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

//92. 反转链表 II
//                    m                    n
//         first     secend
// Head -> head_1 -> head_2 -> head_3 -> head_4 -> head_5 -> Tail
//          prev      cur
//
//         first     secend
// Head -> head_1 <->head_2 <- head_3 <- head_4   head_5 -> Tail
//                                        prev      cur
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	var prev *ListNode
	//定位到m和n
	for m > 1 {
		prev = cur
		cur = cur.Next
		m, n = m-1, n-1
	}
	//反转
	first, secend := prev, cur
	for n > 0 {
		third := cur.Next
		cur.Next = prev
		prev = cur
		cur = third
		n -= 1
	}
	//调整
	if first != nil {
		first.Next = prev
	}else {
		head = prev
	}
	secend.Next = cur
	return head
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
