package main

//845. 数组中的最长山脉   36
func longestMountain(A []int) int {
	length := len(A)
	if length < 3 {
		return 0
	}
	ans, base := 0 , 0
	for base < length{
		end := base
		if end + 1 < length && A[end] < A[end + 1] {
			for end + 1 < length && A[end] < A[end + 1] {
				end++
			}

			if end + 1 < length && A[end] > A[end + 1] {
				for end + 1 < length && A[end] > A[end + 1] {
					end++
				}
				ans = maxs2(ans, end - base + 1)    // int(math.Max(float1,float2))    math.Max 要求是两个float
			}
		}
		base = maxs2(end , base + 1 )
	}
	return ans
}

func maxs2(a , b int) int {
	if a > b {
		return a
	}
	return b
}

//1052. 爱生气的书店老板
func maxSatisfied(customers []int, grumpy []int, X int) int {
	sum, lost := 0,0
	for i := 0 ; i < len(customers) ; i++ {
		sum += customers[i]
		if grumpy[i] == 1 {
			lost += customers[i]
		}else {
			customers[i] = 0
		}
	}
	max , left , clam := 0,0,0
	for r := 0 ; r < len(customers) ; i++ {
		if r - left + 1 <= X{
			clam += customers[r]
		}else {
			clam += customers[r]
			clam -= customers[left]
			left++
		}
		if clam > max {
			max = clam
		}
	}
	return sum - lost + clam
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

//86. 分隔链表
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