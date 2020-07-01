package main

import (
	"fmt"
	"math"
)
//235. 二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}else {
		return root
	}
}

//236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

//二叉树的最大深度
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	Left_hight := maxDepth(root.Left)
	Right_hight := maxDepth(root.Right)
	if Left_hight > Right_hight {
		return Left_hight + 1
	}
	return Right_hight + 1
}

//验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	return isvalid(root,  math.MinInt64, math.MaxInt64)
}

func isvalid(root *TreeNode, min int , max int ) bool {
	if root == nil {
		return true
	}

	if root.Val >= max {
		return false
	}

	if root.Val <= min {
		return false
	}

	return isvalid(root.Left,min , root.Val) && isvalid(root.Right, root.Val , max)
}


//102. 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	return dft(root,0,[][]int{})
}

func dft(root *TreeNode,level int , res [][]int) [][]int {
	if root == nil {
		return res
	}
	if len(res) <= level {
		res = append(res,[]int{root.Val})
	}else{
		res[level] = append(res[level],root.Val)
		fmt.Println(res[level])
	}
	res = dft(root.Left,level + 1, res)
	res = dft(root.Right,level + 1, res)
	return res
}

//105. 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	var root int
	for k , v := range inorder {
		if v == preorder[0] {
			root = k
			break
		}
	}

	return &TreeNode{
		Val : preorder[0],
		Left : buildTree(preorder[1:root+1],inorder[0:root]),
		Right : buildTree(preorder[root+1:],inorder[root+1:]),
	}
}