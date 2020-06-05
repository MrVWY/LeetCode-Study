package main

// 234 回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	//到达链表中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//反转链表
	secend := reverseLinked(slow)
	first := head
	flag := true
	//比较节点
	for flag && secend != nil {
		if first.Val != secend.Val {
			flag = false
		}
		first = first.Next
		secend = secend.Next
	}
	return flag
}

func reverseLinked(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		mid := cur.Next
		cur.Next = prev
		prev = cur
		cur = mid
	}
	return prev
}

// 86 分隔链表  2个空节点+判断+合并
func partition(head *ListNode, x int) *ListNode {
	beforeNode := &ListNode{}
	afterNode := &ListNode{}
	before , after := beforeNode , afterNode
	if head == nil || head.Next == nil {
		return head
	}
	for head !=nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		}else{
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterNode.Next
	return beforeNode.Next
}

// 61 旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil{
		return nil
	}
	if head.Next == nil {
		return head
	}
	count := 1
	node := head
	for node.Next != nil {
		node = node.Next
		count++
	}
	node.Next = head //首尾相连
	node = head //把node指向head节点
	for i := 1 ; i < count-k%count ; i++ {
		node = node.Next
	}
	head, node.Next =node.Next ,nil
	return head
}