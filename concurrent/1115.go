package main

import (
	"fmt"
	"time"
)

type foobar struct {
	n int
	f chan int
	b chan int
}

func (f *foobar) foo() {
	for i := 0 ; i < f.n ; i++ {
		<- f.f
		fmt.Println("foo")
		f.b <- 1
	}
}

func (f *foobar) bar() {
	for i := 0 ; i < f.n ; i++ {
		<-f.b
		fmt.Println("bar")
		f.f <- 1
	}
}

func main() {
	F := new(foobar)
	F.n = 2
	F.f = make(chan int, 1)
	F.b = make(chan int, 1)
	F.f <- 1
	go F.foo()
	go F.bar()
	time.Sleep(4 * time.Second)
}