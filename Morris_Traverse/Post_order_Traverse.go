package Morris_Traverse

import "fmt"

//后序遍历 左-右-跟

func PostOrder(root *TreeNode)  {
	cur := root
	for cur != nil {
		pre := cur.Left
		if pre != nil {
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = cur
				cur = cur.Left
				continue
			}else {
				pre.Right = nil
				printAndRevers(cur.Left)
			}
		}
		cur = cur.Right
	}
	printAndRevers(root)
}

func printAndRevers(node *TreeNode) {
	tail := revers(node)
	cur := tail
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Right
	}
	revers(tail)
}

func revers(node *TreeNode) *TreeNode {
	pre := &TreeNode{}
	next := &TreeNode{}
	for node != nil {
		next = node.Right
		node.Right = pre
		pre = node
		node = next
	}
	return pre
}