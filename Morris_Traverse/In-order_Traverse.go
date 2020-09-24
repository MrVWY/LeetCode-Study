package Morris_Traverse
//中序遍历  左-中-右
import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func In_order(root *TreeNode)  {
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
			}
		}
		fmt.Println(cur.Val)
		cur = cur.Right
	}
}