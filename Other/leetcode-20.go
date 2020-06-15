package main

import (
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

//1299. 将每个元素替换为右侧最大元素
//逆序遍历
func replaceElements(arr []int) []int {
	var ans int = -1
	for i:=len(arr)-1;i>-1;i-- {
		arr[i],ans = ans, max(ans,arr[i])
	}
	return arr
}

func max(a,b int) int {
	if a < b {
		return b
	}
	return a
}


//20. 有效的括号 利用栈
func isValid2(s string) bool {
	stack := make([]int32,len(s))
	stacklength := 0

	for _, v := range s {
		if (v == '(') || (v == '[') ||  (v == '{') {
			stack[stacklength] = v
			stacklength++
		}else{
			if stacklength == 0 {
				return false
			}
			if (v == ')' && stack[stacklength -1] == '(' ) || (v == ']'  && stack[stacklength -1] == '[') || (v == '}'  && stack[stacklength -1] == '{') {
				stacklength--
			}else {
				return false
			}
		}
	}
	return stacklength == 0
}