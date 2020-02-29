package main

import "fmt"

func Sprial_Matrix_One(matrix [][]int) []int {
	if len(matrix) < 1 {
		return []int{}
	}
	left := 0
	right := len(matrix[0]) - 1
	up := 0
	down := len(matrix) - 1
	var r []int
	for left <= right && up <= down {
		//左
		for i := left ; i <= right ; i++ {
			r = append(r, matrix[up][i])
		}
		//下
		for j := up + 1 ; j <= down ; j++ {
			r = append(r, matrix[j][right])
		}
		if left < right && up < down{
			//右
			for i := right ; i >= left ; i-- {
				r = append(r, matrix[down][i])
			}
			//上
			for j := down; j >= up + 1 ; j-- {
				r = append(r, matrix[j][left])
			}
		}
		left++
		right--
		up++
		down--
	}
	return r
}

func Sprial_Matrix_Two(matrix int) [][]int {
	if matrix < 1 {
		return [][]int{}
	}
	left := 0
	right := matrix - 1
	up := 0
	down := matrix - 1
	k := 1
	r := make([][]int,matrix)
	for i:=0 ; i<matrix ; i++ {
		r[i] = make([]int,matrix)
	}
	for left <= right && up <= down {
		//左
		for i := left ; i <= right ; i++ {
			r[up][i] = k
			k++
		}
		//下
		for j := up + 1 ; j <= down ; j++ {
			r[j][right] = k
			k++
		}
		if left < right && up < down{
			//右
			for i := right ; i >= left ; i-- {
				r[down][i] = k
				k++
			}
			//上
			for j := down; j >= up + 1 ; j-- {
				r[j][left] = k
				k++
			}
		}
		left++
		right--
		up++
		down--
	}
	return r
}

func main() {
	a := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
		{10,11,12},
	}
	r := Sprial_Matrix_One(a)
	fmt.Println(r)

	e := Sprial_Matrix_Two(3)
	fmt.Println(e)
}