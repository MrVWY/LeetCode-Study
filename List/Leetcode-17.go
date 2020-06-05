package main

//反转链表 206
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for {
		if cur == nil {
			break
		}
		middler := cur.Next
		cur.Next = pre
		pre = cur
		cur = middler
	}
	return pre
}

//  合并两个有序链表 21
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	pre := head  // 构建虚拟头节点
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			pre.Next = l2
			l2 = l2.Next
		}else{
			pre.Next = l1
			l1 = l1.Next
		}
		pre = pre.Next
	}

	if l1 == nil {
		pre.Next = l2
	}else if l2 == nil {
		pre.Next = l1
	}
	return head.Next
}

//回文链表 (快慢指针+翻转)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head
	var preNode *ListNode = nil
	var preNode2 *ListNode = nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		//从中点开始翻转链表
		preNode2 = slow.Next
		slow.Next = preNode
		preNode = slow
		slow = preNode2
	}

	if fast != nil {
		slow = slow.Next
	}
	res := true
	for slow != nil  {
		if slow.Val != preNode.Val {
			res = false
			break
		}
		slow = slow.Next
		preNode = preNode.Next
	}
	return res
}


//24. 两两交换链表中的节点
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	firtNode := head
	secondNode := head.Next

	firtNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firtNode

	return secondNode
}