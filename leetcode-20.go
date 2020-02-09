package main

import (
	"fmt"
	"strconv"
)

//Fizz Buzz
func fizzBuzz(n int) []string {
	res := []string{}  //make([]string, 0, n)
	for i := 1 ; i <= n ; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			res = append(res,"FizzBuzz")
			continue
		}else if i % 3 == 0 {
			res = append(res,"Fizz")
			continue
		}else if i % 5 == 0 {
			res = append(res,"Buzz")
			continue
		}else {
			number := strconv.Itoa(i)
			res = append(res,number)
		}
	}
	return res
}

//只出现一次的数字 (XOR操作)
func singleNumber(nums []int) int {
	ret := 0
	for _, v := range nums {
		ret ^= v
	}
	return ret
}

//137. 只出现一次的数字 II
//1 &^ 1 = 0
//1 &^ 0 = 1
//0 &^ 1 = 0
//0 &^ 0 = 0
func singleNumber2(nums []int) int {
	a,b := 0,0
	for _,v := range nums{
		a = a ^ v &^ b
		b = b ^ v &^ a
		fmt.Println(a,b)
	}
	return a
}

//832. 翻转图像
func flipAndInvertImage(A [][]int) [][]int {
	for _ , row := range A {
		end := len(row)
		for i := 0 ; i < (end +1)/2 ; i++ {
			row[i] , row[end-1-i] = row[end-1-i] ^ 1 , row[i] ^ 1
		}
	}
	return A
	//for i:=range A{
	//	end:=len(A[i])-1
	//	for j:=0;j<=end;j++{
	//		if A[i][j]==A[i][end]{
	//			A[i][j]=1-A[i][j]
	//			A[i][end]= A[i][j]
	//		}
	//		end--
	//	}
	//}
	//return A
}