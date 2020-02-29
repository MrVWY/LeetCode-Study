package Strings

import (
	"fmt"
	"strconv"
)

//43. 字符串相乘
func multiply(num1 string, num2 string) string {
	var result string
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	length_All := len(num1) + len(num2)
	res := make([]int,length_All)
	for i := len(num1) - 1 ; i >= 0 ; i-- {
		n1, _ := strconv.Atoi(string(num1[i]))
		for j := len(num2) - 1 ; j >= 0 ; j-- {
			n2, _ := strconv.Atoi(string(num2[j]))
			sum := res[i+j+1] + n1 * n2
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}

	}
	fmt.Println(res)
	for i := 0 ; i < len(res) ; i++ {
		if i == 0 && res[i] == 0 {continue}
		result += strconv.Itoa(res[i])
	}
	return result
}