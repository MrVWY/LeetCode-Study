package main

import (
	"fmt"
	"math"
)

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


//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
func lengthOfLongestSubstring(s string) int {
	ans,start,end := 0,0,0
	ma := make(map[byte]byte)
	leng := len(s)
	for start < leng && end < leng {
		Value := s[end]
		if _, ok := ma[Value] ; ok {
			delete(ma , s[start])
			start++
		}else {
			ma[s[end]] = s[end]
			end++
			ans = int(math.Max(float64(ans),float64(end-start)))
			fmt.Println(ans,end,start)
			//            1   1   0
			//            2   2   0
			//            3   3   0
			//            3   4   1
			//            3   5   2
			//            3   6   3
			//            3   7   5
		}
	}
	return ans
}


func main() {
	b := &ListNode{Val: 9, Next:nil}
	a := &ListNode{Val: 9, Next:b}
	d := &ListNode{Val: 1, Next:nil}

	q := addTwoNumbers(a,d)
	fmt.Println(q,q.Next,q.Next.Next)
}