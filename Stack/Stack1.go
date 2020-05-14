package Stack

import "math"

type MinStack struct {
	Stack []int
	Min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Stack : make([]int,0),
		Min : []int{math.MaxInt64},
	}
}


func (this *MinStack) Push(x int)  {
	this.Stack = append(this.Stack, x)
	this.Min = append(this.Min, min(x,this.Min[len(this.Min)-1]))
}


func (this *MinStack) Pop()  {
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.Min = this.Min[:len(this.Min)-1]
}


func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}

func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */