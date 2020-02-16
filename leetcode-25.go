package main

import (
	"fmt"
	"runtime"
)

//1116. 打印零与奇偶数   0-1-0-2-0-3-0-4-0-5
//参考：https://zhuanlan.zhihu.com/p/103712011
type ZeroEvenOdd struct {
	Number           int
	StreamZeroToEven chan struct{}
	StreamZeroToOdd  chan struct{}
	StreamZeroToEnd  chan struct{}
	StreamEvenToZero chan struct{}
	StreamOddToZero  chan struct{}
}

func (o *ZeroEvenOdd) Zero(printNumber func(int)) {
	for i := 0 ; i < o.Number ; i++ {
		select {
		case <- o.StreamOddToZero:
			printNumber(0)
			o.StreamZeroToEven <- struct{}{}
		case <- o.StreamEvenToZero:
			printNumber(0)
			o.StreamZeroToOdd <- struct{}{}
		default:
			runtime.Gosched()
			i--
		}
	}
	//判断当运行到n(即当打印到最后一个number)时,等待SteamEvenToZero 或者 StreamOddToZero 传来结束信号
	if o.Number % 2 == 0 {
		<- o.StreamEvenToZero
	}else {
		<- o.StreamOddToZero
	}
	//待SteamEvenToZero 或者 StreamOddToZero 后, 自己再结束
	o.StreamZeroToEnd <- struct{}{}
}

func (o *ZeroEvenOdd) Even(printNumber func(int)) {
	number := o.Number - o.Number % 2 //计算
	for i := 2 ; i <= number; {
		<- o.StreamZeroToEven //等待Zero的通知
		printNumber(i)
		i += 2
		o.StreamEvenToZero <- struct{}{} //通知Zero,已经打印Even
	}
}

func (o *ZeroEvenOdd) Odd(printNumber func(int)) {
	number := (o.Number + 1) - ((o.Number + 1) % 2 + 1)
	for i := 1 ; i <= number; i += 2{
		<- o.StreamZeroToOdd //等待Zero的通知
		printNumber(i)
		o.StreamOddToZero <- struct{}{} //通知Zero,已经打印Odd
	}
}

//打印函数
func printNumber(i int)  {
	fmt.Printf("%d", i)
}

func main(){
	var StartZeroEvenOdd = func(number int) {
		var zero = &ZeroEvenOdd{
			Number:           number,
			StreamZeroToEven: make(chan struct{}),
			StreamZeroToOdd:  make(chan struct{}),
			StreamZeroToEnd:  make(chan struct{}),
			StreamEvenToZero: make(chan struct{}),
			StreamOddToZero:  make(chan struct{}),
		}
		go func() { zero.StreamEvenToZero <- struct{}{} }()
		go zero.Zero(printNumber)
		go zero.Even(printNumber)
		go zero.Odd(printNumber)
		<- zero.StreamZeroToEnd
		fmt.Println()
	}

	for k := range [14]int{}{
		fmt.Printf("Case %2d: ", k)
		StartZeroEvenOdd(k)
	}
}