package main

import (
	"fmt"
	"runtime"
	"sync"
)

type FizzBuzz struct {
	n int
	wg *sync.WaitGroup
	streamBaton chan int
}

func (f *FizzBuzz) Prion(pass func(int) bool, prins func(int)) {
	defer f.wg.Done()
	for i := 0 ; i < f.n ; i++ {
		if pass(i) {
			nextnum := <- f.streamBaton //接
			if i == nextnum {
				prins(i)
				f.streamBaton <- i+1 //交
			}else {
				f.streamBaton <- nextnum //还
				i--
			}
			runtime.Gosched()
		}
	}
}

func (f *FizzBuzz) fizz() {
	pass := func(i int) bool {
		return (0 == i % 3 ) && (0 != i % 5 )
	}
	prin := func(i int) {
		fmt.Printf("Fizz(%d) ", i)
	}
	f.Prion(pass, prin)
}

func (f *FizzBuzz) buzz() {
	pass := func(i int) bool {
		return (0 != i % 3 ) && (0 == i % 5 )
	}
	prin := func(i int) {
		fmt.Printf("Buzz(%d) ", i)
	}
	f.Prion(pass, prin)
}

func (f *FizzBuzz) fizzbuzz() {
	pass := func(i int) bool {
		return 0 == i % (3*5)
	}
	prin := func(i int) {
		fmt.Printf("FizzBuzz(%d) ", i)
	}
	f.Prion(pass, prin)
}

func (f *FizzBuzz) printnum() {
	pass := func(i int) bool {
		return (0 != i % 3 ) && (0 != i % 5 )
	}
	prin := func(i int) {
		fmt.Printf("%d ", i)
	}
	f.Prion(pass, prin)
}

func main() {

	for num := 0 ; num <= 20 ; num++ {
		f := &FizzBuzz{
			num,
			&sync.WaitGroup{},
			make(chan int, 1),
		}
		f.wg.Add(4)
		go f.fizz()
		go f.buzz()
		go f.fizzbuzz()
		go f.printnum()
		f.streamBaton <- 0
		f.wg.Wait()
		close(f.streamBaton)
		fmt.Println()
	}

}