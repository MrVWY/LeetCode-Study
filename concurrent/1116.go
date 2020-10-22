package main

//1165
import (
	"fmt"
	"runtime"
)

func printNumber(n int) {
	fmt.Printf("%d", n)
}

type ZeroEvenOdd struct {
	n int
	streamEvenToZero chan struct{}
	streamOddToZero chan struct{}
	streamZeroToEven chan struct{}
	streamZeroToOdd chan struct{}
	streamZeroToEnd chan struct{}
}

func (Z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for i := 0 ; i < Z.n ; i++ {
		select {
		case <-Z.streamOddToZero:
			printNumber(0)
			Z.streamZeroToEven <- struct{}{}
		case <-Z.streamEvenToZero:
			printNumber(0)
			Z.streamZeroToOdd <- struct{}{}
		default:
			runtime.Gosched()
			i--
		}
	}
	if 0 == Z.n % 2 {
		<- Z.streamEvenToZero
	}else {
		<- Z.streamOddToZero
	}

	Z.streamZeroToEnd <- struct{}{}
}

func (Z *ZeroEvenOdd) Odd(printNumber func(int)) {
	oddUpper := ( (Z.n + 1) - (Z.n + 1) % 2 ) -1
	for i := 1 ; i <= oddUpper ; i += 2{
		<- Z.streamZeroToOdd
		printNumber(i)
		Z.streamOddToZero <- struct{}{}
	}
}

func (Z *ZeroEvenOdd) Even(printNumber func(int)) {
	evenUpper := Z.n - Z.n % 2
	for i := 2 ; i <= evenUpper ; {
		<- Z.streamZeroToEven
		printNumber(i)
		i += 2
		Z.streamEvenToZero <- struct{}{}
	}
}

func main() {
	var PrintZeroEvenOdd = func(num int) {
		var zeo = &ZeroEvenOdd{
			n : num,
			streamOddToZero: make(chan struct{}),
			streamZeroToOdd: make(chan struct{}),
			streamEvenToZero: make(chan struct{}),
			streamZeroToEven: make(chan struct{}),
			streamZeroToEnd: make(chan struct{}),
		}

		go func() { zeo.streamEvenToZero <- struct{}{} }()
		go zeo.Zero(printNumber)
		go zeo.Even(printNumber)
		go zeo.Odd(printNumber)
		<- zeo.streamZeroToEnd
		fmt.Println()
	}
	for num := range [14]int{} {
		fmt.Printf("Case %2d:", num)
		PrintZeroEvenOdd(num)
	}
}