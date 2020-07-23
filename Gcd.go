package main

import "fmt"

func main() {
	var a , b int

	fmt.Println(a*b/gcd(a,b)) //最小公倍数 = 两数相乘/ 最大公约数
}

//最大公约数
func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a % b == 0 {
		return b
	}
	return gcd(b, a%b)
}