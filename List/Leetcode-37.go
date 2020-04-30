package main

import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := []int{},[]int{}
	for l1 != nil {
		stack1 = append(stack1,l1.Val)
		l1 = l1.Next
	}

	for l2 != nil{
		stack2 = append(stack2,l2.Val)
		l2 = l2.Next
	}
	carry := 0
	res := &ListNode{}
	i,j := len(stack1), len(stack2)
	for i != 0 || j != 0 || carry != 0 {
		temp1 ,temp2 := 0,0
		if i != 0 {
			i = i - 1
			temp1 = stack1[i]
		}

		if j != 0 {
			j = j - 1
			temp2 = stack2[j]
		}
		res.Next = &ListNode{(temp1 + temp2 + carry)%10, res.Next}
		carry = (temp1 + temp2 + carry) / 10
		//node.Next = res.Next
		//res.Next = node
	}
	return res.Next
}

//23. 合并K个排序链表  优先队列
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Queue []*ListNode
func (q Queue) Len() int {return len(q)}
func (q Queue) Less(i, j int) bool {
	return q[i].Val < q[j].Val
}
func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	*q = old[0:n-1]
	return item
}

func (q *Queue) Push(x interface{}) {
	item := x.(*ListNode)
	*q = append(*q, item)
}

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{Val:-1,Next:nil}
	t := head
	if len(lists) == 0 {
		return head.Next
	}

	queue := make(Queue,0)
	for i := 0 ; i < len(lists) ; i++ {
		if lists[i] != nil {
			queue = append(queue,lists[i])
		}
	}
	heap.Init(&queue)
	for len(queue) != 0 {
		item := heap.Pop(&queue).(*ListNode)
		next := item.Next
		item.Next = t.Next
		t.Next = item
		t = item

		if next != nil {
			heap.Push(&queue,next)
		}
	}

	return head.Next
}