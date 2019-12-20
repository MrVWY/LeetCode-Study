package main

import "fmt"

//杨辉三角
func generate(numRows int) [][]int {
	v := make([][]int,numRows)
	for i := 0 ; i < numRows ; i++ {
		v[i] = make([]int,i + 1)
	}
	for j := 0 ; j < numRows ; j++ {
		for k := 0 ; k <= len(v[j]) - 1 ; k++ {
			fmt.Println(len(v[j]))
			if k == 0 || k == len(v[j]) - 1 {
				v[j][k] = 1
			}else {
				v[j][k] = v[j-1][k-1] + v[j-1][k]
			}
		}
	}
	return v
}