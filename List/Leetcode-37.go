package main

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