package main

import "math"

//Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//101. 对称二叉树
func isSymmetric(root *TreeNode) bool {
	return isvalid2(root,root)
}

func isvalid2(root1 *TreeNode,root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return (root1.Val == root2.Val) && isvalid2(root1.Left,root2.Right) && isvalid2(root1.Right,root2.Left)
}

//617. 合并二叉树
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)

	return t1
}

//226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

//965. 单值二叉树
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	val := root.Val
	return compare(root.Left, val) && compare(root.Right, val)
}

func compare(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val == val{
		return compare(root.Left, val) && compare(root.Right, val)
	}
	return false
}

//110. 平衡二叉树
func isBalanced(root *TreeNode) bool {
	return root == nil || isBalanced(root.Left) && math.Abs(float64(height(root.Right)-height(root.Left))) < 2 && isBalanced(root.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l := height(node.Left)
	r := height(node.Right)
	if r > l {
		return r+1
	}else{
		return l+1
	}
}