package Breadth_First_Search

//从上到下打印二叉树
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) []int {
	res := make([]int,0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode,0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

	}

	return res
}

//从上到下打印二叉树II
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	p := []*TreeNode{root}
	for i := 0 ; len(p) > 0 ; i++  {
		result = append(result, []int{})
		q := []*TreeNode{}
		for j := 0 ; j < len(p); j++ {
			node := p[j]
			result[i] = append(result[i],node.Val)
			if node.Left != nil {
				q = append(q,node.Left)
			}
			if node.Right != nil {
				q = append(q,node.Right)
			}
		}
		p = q
	}
	return result
}

//从上到下打印二叉树III
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	p := []*TreeNode{root}
	height := 0
	for i := 0 ; len(p) > 0 ; i++  {
		temp := make([]int,0)
		q := []*TreeNode{}
		for j := 0 ; j < len(p); j++ {
			node := p[j]
			if node.Left != nil {
				q = append(q,node.Left)
			}
			if node.Right != nil {
				q = append(q,node.Right)
			}
			temp = append(temp, node.Val)
		}
		if height % 2 == 0 {
			result = append(result, temp)
		}else{
			ttmp := []int{}
			for k := len(temp)-1; k >= 0; k-- {
				ttmp = append(ttmp,temp[k])
			}
			result = append(result, ttmp)
		}
		height++
		p = q
	}
	return result
}