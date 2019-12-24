package main

import (
	"fmt"
	"strconv"
	"strings"
)
//字符串转换整数 (atoi)
func myAtoi(str string) int {
	res := ""
	str = strings.TrimSpace(str)
	if len(str) == 0  {
		fmt.Println("1")
		return 0
	}else if len(str) == 1 {
		if (str[0] >= 48  && str[0] < 58) {
			res += str[0:1]
			a, _ := strconv.Atoi(res)
			return a
		}else {
			return 0
		}
	}
	i := a(strings.TrimSpace(str))
	return i
}

func a(str string) int {
	res := ""
	flag := 1
	if (str[0] >= 48  && str[0] < 58) || str[0] == 45  || str[0] == 32 || str[0] == 43 {
		fmt.Println("2")
		for i := 0 ; i < len(str) ; i++ {
			// if str[i] == 32 && i == 0  {
			//     continue
			// }else
			if str[i] == 45 && i == 0  {
				flag = -1
				continue
			}else if str[i] == 43 && i == 0 {
				flag = 1
				continue
			}else if str[i] > 57 || str[i]  < 48 {
				break
			}
			res += str[i:i+1]
		}
	}
	a, _ := strconv.Atoi(res)
	if a >= 2147483648 && flag == -1 {
		a = 2147483648
		return  a * flag
	}else if a > 2147483648 {
		a = 2147483648 - 1
		return  a * flag
	}else if a == 2147483648 {
		a = 2147483647
		return  a * flag
	}
	return a * flag
}

//. 整数反转
func reverse(x int) int {
	flag := 1
	if x == 0 {
		return x
	}
	if x < 0 {
		flag = -1
	}
	x = x * flag
	res := 0
	for {
		if x == 0 {
			break
		}
		pop := x % 10
		x = x / 10
		res = res * 10 + pop
		if res >= 2147483648 {
			return 0
		}
	}
	return res * flag
}