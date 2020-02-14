package main

//1047. 删除字符串中的所有相邻重复项   栈的思想
func removeDuplicates2(S string) string {
	stack := []byte{}
	for i := 0 ; i < len(S) ; i++ {
		if len(stack) == 0 {
			stack = append(stack,S[i])
		}else {
			if S[i] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}else{
				stack = append(stack,S[i])
			}
		}
	}
	return string(stack)
}

//面试题28. 对称的二叉树
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return h(root.Left,root.Right)
}

func h(t1 *TreeNode,t2 *TreeNode) bool{
	if t1 == nil && t2 ==nil {
		return true
	}else if t1 == nil || t2 == nil {
		return false
	}

	return t1.Val == t2.Val && h(t1.Left,t2.Right) && h(t1.Right,t2.Left)
}

//面试题59 - I. 滑动窗口的最大值
func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	if len(nums) == 0 {
		return res
	}else if len(nums) < k {
		res = append(res,max2s(nums))
		return res
	}
	for i := 0 ; i < len(nums) - k + 1 ;i++ {
		res = append(res,max2s(nums[i:i+k]))
	}
	return res
}

func max2s(nums []int) int {
	max := nums[0]
	for i := 1 ; i < len(nums) ; i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}