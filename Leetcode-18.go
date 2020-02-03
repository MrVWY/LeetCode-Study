package main

import (
	"fmt"
	"math"
)

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

//对称二叉树
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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


//二叉树的层次遍历
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