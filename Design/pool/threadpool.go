package main

import (
	"fmt"
	"time"
)

//define a type of task , named Task
type Task struct {
	f func() error
}

//create a mission of Task
func NewTask(g_f func() error) *Task {
	t := Task{f:g_f}
	return &t
}

//Task need a function about perform business
func (t *Task) Execute() {
	_ = t.f()
}

//define a type of pool
type Pool struct {
	//Entrance to the outside
	EntryChannel chan *Task

	//External port
	JobsChannel chan *Task

	//max worker number
	WorkerNum int
}

//create a pool
func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		JobsChannel: make(chan *Task) ,
		WorkerNum:    cap,
	}
	return &p
}

//Pool need a function about perform business
//create worker
func (p *Pool) worker(workerId int) {
	//get mission in JobsChannel
	//select {
	//case task := <-p.JobsChannel:
	//	task.Execute()
	//	fmt.Printf("%d 已经执行完工作", workerId)
	//}

	for task := range  p.JobsChannel {
		task.Execute()
		fmt.Printf("%d 已经执行完工作", workerId)
	}
}

func (p *Pool) run() {
	//create worker base on WorkerNum
	for i := 0 ; i < p.WorkerNum ; i++{
		go p.worker(i)
	}
	//get mission in EntryChannel
	//select {
	//case task := <- p.EntryChannel:
	//	p.JobsChannel <- task
	//}

	for task := range  p.EntryChannel {
		p.JobsChannel <- task
	}
}

func main() {
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})

	p := NewPool(4)

	go func() {
		for {
			p.EntryChannel <- t
		}
	}()

	p.run()
}