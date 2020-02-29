package main

//二叉搜索树中的两个节点被错误地交换。
//请在不改变其结构的情况下，恢复这棵树。
//          3
//         / \
//        1   4
// 首先 RecoverTree() -> In_Order_Traversal(root) ->
// In_Order_Traversal(root.Left)[ In_Order_Traversal(root) ]  pre = 1  -> In_Order_Traversal(root)的In_Order_Traversal(root.Left)之后,这时pre = 3 ->
// In_Order_Traversal(root.Right)[ In_Order_Traversal(root) ]  pre = 4
type treeNodes struct {
	Value int
	Left *treeNodes
	Right *treeNodes
}

var pre, first, second *treeNodes

func RecoverTree(Root *treeNodes)  {
	pre = nil
	first = nil
	second = nil
	In_Order_Traversal(Root)
	first.Value, second.Value = second.Value, first.Value
}

func In_Order_Traversal(root *treeNodes)  {
	if root == nil{
		return
	}
	In_Order_Traversal(root.Left)
	if pre != nil && pre.Value > root.Value{
		if first == nil{
			first = pre
			second = root
		}else {
			second = root
		}
	}
	pre = root
	In_Order_Traversal(root.Right)
}