package Design

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