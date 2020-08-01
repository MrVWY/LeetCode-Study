package main

import "fmt"

//二进制转十进制
func BinaryToDecimal(binary string) int {
	n := len(binary)
	ans := 0
	for i := 0 ; i < n ; i++ {
		ans = ans * 2 + int(binary[i] - '0')
	}
	return ans
}

func main() {
	fmt.Println(BinaryToDecimal("10101"))
}