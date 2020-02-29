package main

import "fmt"

//132模式 (参考大佬)

func find132pattern(nums []int) bool {
	length := len(nums)
	if length < 3 {
		return false
	}
	min := findTheminI(nums)
	fmt.Println(min)
	stack := newStack(length)
	for i := length - 1 ; i >= 0 ; i-- {
		fmt.Println("1")
		if nums[i] > min[i] {
			fmt.Println("2")
			for !stack.isEmpty() && stack.tops() <= min[i] {
					stack.pop()
			}
			if !stack.isEmpty() && stack.tops() < nums[i] {
				return true
			}
			stack.push(nums[i])
		}
	}
	return false
}

func findTheminI(nums []int) []int {
	min := make([]int,len(nums))
	min[0] = nums[0]
	for i := 1 ; i < len(nums) ; i++ {
		min[i] = CompareMin(nums[i],min[i-1])
	}
	return min
}

func CompareMin(a,b int) int {
	if a >= b {
		return b
	}
	return a
}

type Stacks struct {
	items []int
	top int
}

func newStack(a int) *Stacks {
	return &Stacks{
		items: make([]int,a),
		top:   -1,
	}
}

func (s *Stacks) isEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}
//出栈
func (s *Stacks) pop()  int {
	if s.isEmpty() {
		return 0
	}
	value := s.items[s.top]
	s.top--
	return value
}
//入栈
func (s *Stacks) push(value int) {
	if s.top < 0 {
		s.top = 0
	}else {
		s.top++
	}

	if s.top > len(s.items) - 1 {
		s.items = append(s.items,value)
	}else {
		s.items[s.top] = value
	}
}

func (s *Stacks) tops() int {
	if s.isEmpty() {
		return 0
	}
	return s.items[s.top]
}

func main() {
	//a := []int{-2,1,-1}
	b := []int{2,3,1,2}
	s := find132pattern(b)
	fmt.Println(s)
}