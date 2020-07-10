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

// 61 旋转链表 快慢指针
//A={1,3,5,7,9,11}
//B={2,4,9,11}
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

//160. 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		}else {
			curA = curA.Next
		}

		if curB == nil {
			curB = headA
		}else {
			curB = curB.Next
		}
	}

	return curA
}

//分割链表
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	d1, d2 := &ListNode{}, &ListNode{}
	cur1, cur2 := d1,d2
	cur := head
	for cur != nil {
		if cur.Val < x {
			cur1.Next = cur
			cur = cur.Next
			cur1 = cur1.Next
			cur1.Next = nil
		}else if cur.Val >= x {
			cur2.Next = cur
			cur = cur.Next
			cur2 = cur2.Next
			cur2.Next = nil
		}
	}
	cur1.Next = d2.Next
	return d1.Next
}

//链表相交
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	t1, t2 := headA, headB
	for t1 != t2 {
		if t1 == nil {
			t1 = headB
		}else {
			t1 = t1.Next
		}
		if t2 == nil {
			t2 = headA
		}else{
			t2 = t2.Next
		}
	}
	return t1
}

//返回倒数第 k 个节点
func kthToLast(head *ListNode, k int) int {
	p := head
	for i := 0 ; i < k ; i++ {
		p = p.Next
	}

	for p != nil {
		p = p.Next
		head = head.Next
	}
	return head.Val
}

//148. 排序链表
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	Left, Right := sortList(head), sortList(mid)
	return merge(Left, Right)
}

func merge(Left, Right *ListNode) *ListNode {
	res := &ListNode{Val: 0}
	h := res
	for Left != nil && Right != nil {
		if Left.Val <= Right.Val {
			h.Next, Left = Left, Left.Next
		}else{
			h.Next, Right = Right, Right.Next
		}
		h = h.Next
	}
	if Left != nil {
		h.Next, Left = Left, Left.Next
	}
	if Right != nil {
		h.Next, Right = Right, Right.Next
	}
	return res.Next
}