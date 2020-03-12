package main

import "fmt"

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


//515. 在每个树行中找最大值
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//以[1,3,2]  分层！！
//-1 2 {root,nil}
//-1 4 {root,   nil,                 root.Left,root.Right}
//     { 1    层与层的分隔               2           2    }
//1	   {root,nil,root.Left,root.Right,nil}
//1 5
//1 5
//1 5
//4	   {root,nil,root.Left,root.Right,nil,nil}
func largestValues(root *TreeNode) []int {
	ret := []int{}
	if root == nil{
		return ret
	}
	array := []*TreeNode{root,nil}
	last := -1
	for {
		if array[last+1] == nil {
			break
		}
		max := array[last+1].Val
		for i := last+1 ; i < len(array) ; i++ {
			fmt.Println(last,len(array)) //一个fmt提交上去用时+2ms................
			if array[i] == nil { //root的左右节点先加进array,在遇到第i层的分隔符，当第i+1层的节点都加进来以后再加上第i+1层的分隔符
				array = append(array, nil)
				last = i //换层也就是换index
				fmt.Println(last)
				break
			}
			if array[i].Left != nil{
				array = append(array, array[i].Left)
			}
			if array[i].Right != nil{
				array = append(array, array[i].Right)
			}
			if array[i].Val > max{
				max = array[i].Val
			}
		}
		ret = append(ret, max)
	}
	return ret
}

//以[1,3,2]  分层！！
//result = [第一层，第二层 ]
func largestValues(root *TreeNode) []int {
	var (
		result []int
		dfs func(node *TreeNode,depath int)
	)
	dfs = func(node *TreeNode,depath int) {
		if node == nil {
			return
		}

		if depath > len(result) {
			result = append(result,node.Val)
		}else if node.Val > result[depath-1] {
			result[depath-1] = node.Val
		}

		dfs(node.Left,depath+1)
		dfs(node.Right,depath+1)
	}
	return result
}
