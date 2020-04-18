package main

import (
	"fmt"
)

func heapSort(input []int){
	inputLen := len(input)
	if inputLen == 0 {
		return
	}
	for i:=0; i<inputLen; i++{
		fmt.Println("i",input[i:])
		minAjust(input[i:])
	}
}

func minAjust(input []int){
	inputLen := len(input)
	if inputLen <= 1{
		return
	}
	for i:= inputLen/2 -1; i>=0; i--{
		if (2*i+1 <= inputLen-1) && (input[i] >= input[2*i+1]){
			input[i], input[2*i+1] = input[2*i+1], input[i]
		}
		if (2*i+2<= inputLen-1) && (input[i] >= input[2*i+2]){
			input[i], input[2*i+2] = input[2*i+2], input[i]
		}
		fmt.Println(input)
	}
}

func main (){
	testSlice := []int {10,5,4,2,3,1}
	heapSort(testSlice)
	fmt.Println(testSlice)
}