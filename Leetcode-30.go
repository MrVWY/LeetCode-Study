package main

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

//405. 数字转换为十六进制数
//首先要知道16进制数在计算机中是怎么表示的，由于对于16进制一个有16个数（0,1,2,3,4,5,6,7,8,9,a,b,c,d,e,f）所以我们只需要使用4位 2^4=16刚好可以表示16个数，所以计算机在内部是用4位来表示16进制的0000表示0,0001表示1.....1111表示f.
//所有题目给我们一个int型的数字num,一个int= 4字节 = 32位，所以，这个32位数字可以被分隔成8个4位数，
//所以我们只需要从头到尾每4位每4的的遍历，每遍历一个4位就将其转化为16进制的数即可，
//所以每4位遍历就变成了0001,1010对于0001 对应16进制中的1,1010对应16进制中的a所以结果就是为1a.
//所以对于此题我们只需要从尾部开始用位运算取出每一个4位，再转化为16进制即可只需要注意前面为0的16进制应该删除即可。首先需要设一个数字mask = 0xf也就是右边4位为1其余均为0，
//此时我们把数字num与之进行与运算就可以取出，num右边4位的值，然后再把num向右边移动4位，继续&mask，重复操作
func toHex(num int) string {
	const (
		hexStr = "0123456789abcdef"
		mask   = 0xf
	)
	res := make([]byte,8)
	index := 7
	for i := 7 ; i >= 0 ; i-- {
		val := num & mask //获得num后四位二进制数
		res[i] = hexStr[val]
		if val > 0 {
			index = i
		}
		num >>= 4 //将num右移4位
	}
	return string(res[index:])
}

//232. 用栈实现队列
type MyQueue struct {
	item []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		item : make([]int,0),
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.item = append(this.item,x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}
	Value := this.item[0]
	this.item = this.item[1:]
	return Value
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	return this.item[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if this.item == nil || len(this.item) == 0 {
		return true
	}
	return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */