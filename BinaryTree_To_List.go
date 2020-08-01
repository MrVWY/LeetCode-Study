package main

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

//114. 二叉树展开为链表
//（1）前序遍历
//（2）寻找前驱节点
func flatten(root *TreeNode)  {
	cur := root
	//寻找前驱节点
	for cur != nil {
		if cur.Left != nil {
			nex := cur.Left
			pre := nex
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur.Right
			cur.Left, cur.Right = nil, nex
		}
		cur = cur.Right
	}
}