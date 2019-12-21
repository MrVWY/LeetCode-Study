package main

import "fmt"

//两数相加
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	carry := 0
	NewNode := &ListNode{}
	cur := NewNode
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		Val := (sum + carry) % 10
		carry = (sum + carry) / 10
		cur.Next = &ListNode{Val:Val}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{Val:carry}
	}
	return NewNode.Next
}

func main() {
	b := &ListNode{Val: 9, Next:nil}
	a := &ListNode{Val: 9, Next:b}
	d := &ListNode{Val: 1, Next:nil}

	q := addTwoNumbers(a,d)
	fmt.Println(q,q.Next,q.Next.Next)
}