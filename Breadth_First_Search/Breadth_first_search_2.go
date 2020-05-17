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


//从上到下打印二叉树III