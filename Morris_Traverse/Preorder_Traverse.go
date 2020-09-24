package Morris_Traverse

import "fmt"

//前序遍历  跟-左-右

func Preorder_Traverse(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			pre := cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = cur
				fmt.Println(cur.Val) //打印根节点
				cur = cur.Left
				continue
			}else {
				pre.Right = nil
			}
		}else {
			fmt.Println(cur.Val)
		}
		cur = cur.Right
	}
}