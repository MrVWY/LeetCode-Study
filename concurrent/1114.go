package main

import "fmt"

type foo struct {
	s chan bool
	t chan bool
	o chan bool
}

func (f *foo) first() {
	fmt.Println("one")
	f.s <- true
}

func (f *foo) second() {
	<- f.s
	fmt.Println("second")
	f.t <- true
}

func (f *foo) third() {
	<- f.t
	fmt.Println("third")
	f.o <- true
}

func main() {
	f := foo{
		make(chan bool),
		make(chan bool),
		make(chan bool),
	}
	go f.first()
	go f.second()
	go f.third()
	<- f.o
}