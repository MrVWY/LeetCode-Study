package Reverse_List

type ListNode struct {
	Val int
	Next *ListNode
}
/*
	注意2种反转链表情况
 */

//206. 反转链表
func reverseList1(head *ListNode) *ListNode {
	var pre *ListNode
	curd := head
	for {
		if curd == nil {
			break
		}
		mid := curd.Next
		curd.Next = pre
		pre = curd
		curd = mid
	}
	return pre
}

//25. K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	heads := &ListNode{Next:head}
	pre := heads
	for {
		tail := pre
		for i := 0 ; i < k ; i++ {
			tail = tail.Next
			if tail == nil {
				return heads.Next
			}
		}
		head, tail = reverseList(head,tail)
		pre.Next = head
		pre = tail
		head = tail.Next
	}
	return heads.Next
}

func reverseList(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}
