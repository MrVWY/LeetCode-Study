package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
//       0
//    1    4
//    2    3

type DiningPhilosophers struct {
	wg                     *sync.WaitGroup
	streamForks            [5]chan interface{}
	missingDoubleForkTimes int
}

func (this *DiningPhilosophers) WantToEat(philosopher int, pickLeftFork func(int), pickRightFork func(int), eat func(int), putLeftFork func(int), putRightFork func(int)) {
	defer this.wg.Done()
	var leftNum = (philosopher + 4 ) % 5
	var rightNum = (philosopher + 6 ) % 5
	for {
		select {
		case this.streamForks[leftNum] <- philosopher:
			PickLeftFork(philosopher) //拿起左边的筷子
			select {
			case this.streamForks[rightNum] <- philosopher:
				PickRightFork(philosopher) //拿起右边的筷子
				Eat(philosopher)
				<- this.streamForks[leftNum] //放下左边的筷子
				PutLeftFork(philosopher)
				<- this.streamForks[rightNum] //放下右边的筷子
				PutRightFork(philosopher)
				return
			default:
				fmt.Printf("Philosopher %d can't pick fork %d.\n", philosopher, rightNum)
				<-this.streamForks[leftNum]
				PutLeftFork(philosopher)
			}
		default:
			fmt.Printf("Philosopher %d can't pick fork %d.\n", philosopher, leftNum)
		}
		this.missingDoubleForkTimes++
		Think()
	}
}

func Eat(philosopher int) {
	fmt.Printf("===== Philosopher %d have eaten. =====\n", philosopher)
}

//思考
func Think() {
	Random := func(max int) int {
		rand.Seed(time.Now().Unix())
		return rand.Int() % (max + 1)
	}
	<-time.After(time.Millisecond * time.Duration(Random(50)))
}

func PickLeftFork(philosopher int) {
	var leftNum = (philosopher + 4) % 5
	fmt.Printf("Philosopher %d picked fork %d.\n", philosopher, leftNum)
}

func PickRightFork(philosopher int) {
	var rightNum = (philosopher + 6) % 5
	fmt.Printf("Philosopher %d picked fork %d.\n", philosopher, rightNum)
}

func PutLeftFork(philosopher int) {
	var leftNum = (philosopher + 4) % 5
	fmt.Printf("Philosopher %d putted fork %d.\n", philosopher, leftNum)

}

func PutRightFork(philosopher int) {
	var rightNum = (philosopher + 6) % 5
	fmt.Printf("Philosopher %d putted fork %d.\n", philosopher, rightNum)

}

func main() {
	diningPhilosophers := DiningPhilosophers{
		wg: &sync.WaitGroup{},
	}

	// Channel 初始化
	for i := range diningPhilosophers.streamForks {
		diningPhilosophers.streamForks[i] = make(chan interface{}, 1)
	}

	// 哲学家开始行动
	start := time.Now()
	for i := range diningPhilosophers.streamForks {
		diningPhilosophers.wg.Add(1)
		go diningPhilosophers.WantToEat(i, PickLeftFork, PickRightFork, Eat, PutLeftFork, PutRightFork)
	}


	diningPhilosophers.wg.Wait()
	fmt.Println("Spent time:", time.Now().Sub(start))
	fmt.Printf("Missing double forks %d times.", diningPhilosophers.missingDoubleForkTimes)
}
