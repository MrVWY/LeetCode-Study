package main

/*
&(and运算、按位与) : 相同位的两个数字都为1,则为1；若有一个不为1，则为0 (任何二进制位和0进行&运算,结果是0; 和1进行&运算结果是原值)
^(xor运算、按位异或) : 任何相同二进制位进行^运算,结果是0; 不相同二进制位^运算结果是1
|(or运算、按位或) :  任何二进制位和0进行|运算,结果是原值; 和1进行|运算结果是1
~(not运算、按位取反) :
<<(左移、按位左移) : 空位补0,被移除的高位丢弃
>>(右移、按位右移) : 被移位的二进制最高位是0,右移后,空缺位补0; 最高位是1, 最高位补1
 */

//位运算判断一个数为奇偶数 (看该数二进制的最后一位是0的话那么就为偶数,是1的话就为奇数)
/*
	5,二进制代码为101
	5 & 1
-> 101 & 001 = 001
*/
func bitwise_operation_is_oddOreven_number(num int) bool {
	return (num & 1) == 0
}


//位运算判断俩个数为异号 (看该数二进制的最高位是0的话那么就为正数,是1的话就为负数)
func bitwise_operation_is(a,b int) bool {
	return (a^b) < 0
}


//位运算加法
/*
（1）0+0=0
（2）1+0=1
（3）0+1=1
（4）1+1=0

	A + B = A^B + (A&B) <<1
	(A&B) <<1 表示A+B的进位
例子：  a:13      b:11
二进制：a:01101	 b:01011
运算 ：
	//位运算操作
	x=num1^num2 : 0 0 1 1 0
	y=num1&num2 : 0 1 0 0 1

	//因为是进位，y要向前移动一位
	y=y<<1      : 1 0 0 1 0
	//将y当成新的数值继续进行位运算
	x2=x^y      : 1 0 1 0 0
	y2=x&y      : 0 0 0 1 0

	y2=y2<<1    : 0 0 1 0 0
	x3=x2^y2    : 1 0 0 0 0
	y3=x2&y2    : 0 0 1 0 0

	y3=y3<<1    : 0 1 0 0 0
	//当y=0时结束
	x4=x3^y3    : 1 1 0 0 0
	y4=x3&y3    : 0 0 0 0 0

    return x4
 */
func bitwise_operation_number_addition(a,b int) int {
	x := a^b
	y := a&b
	for y != 0 {
		y = y << 1
		temp := x //当前x值
		x = x ^ y //运算后x值
		y = temp & y
	}
	return x
}


//位运算减法
/*
	A-B = A + (-B)
	-B = ~B + 1
 */
func bitwise_operation_number_subtraction(a,b int) int {
	b = ^b + 1
	result := bitwise_operation_number_addition(a,b)
	return result
}