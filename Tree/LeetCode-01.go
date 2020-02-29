package main

import "fmt"

//二叉树 求其最大路径和
//                 -10
//                 /  \
//                5   20
//                   /  \
//                  15   7
type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

//计算最大值
func Max(a[] int) int {
	max := -1<<31
	for i := 0;i<len(a);i++  {
		if max < a[i]{
			max = a[i]
		}
	}
	return max
}

func TreeNodeSumRoot(Root *TreeNode) int {
	if Root == nil{
		return  0
	}
	if Root.Left != nil && Root.Right != nil {
		Lr := TreeNodeSonSum(Root.Left)
		Rr := TreeNodeSonSum(Root.Right)
		L := TreeNodeSumRoot(Root.Left) //以Root.Left为根节点
		R := TreeNodeSumRoot(Root.Right)//以Root.Right为根节点
		return Max([]int{Root.Value + Lr, Root.Value + Rr, Root.Value + Lr + Rr, Root.Value, Lr, Rr, L, R})
	} else if Root.Left != nil{
		Lr := TreeNodeSonSum(Root.Left)
		L := TreeNodeSumRoot(Root.Left)
		return Max([]int{Root.Value + Lr, Root.Value, Lr, L})
	} else if Root.Right != nil{
		Rr := TreeNodeSonSum(Root.Right)
		R := TreeNodeSumRoot(Root.Right)
		return Max([]int{Root.Value + Rr, Root.Value, Rr, R})
	}
	return Root.Value
}

func TreeNodeSonSum(Root *TreeNode) int {
	if Root == nil{
		return 0
	}
	L := TreeNodeSonSum(Root.Left)
	R := TreeNodeSonSum(Root.Right)
	return Max([]int{Root.Value + L , Root.Value + R , Root.Value})  // 20 + 15 ， 20 + 7 ， 20
}

func main() {
	Treeson1_2 := &TreeNode{Value: 7, Left:  nil, Right: nil,}
	Treeson1_1 := &TreeNode{Value: 15, Left:  nil, Right: nil,}
	Treeson1 := &TreeNode{Value: 20, Left:  Treeson1_1, Right: Treeson1_2,}
	Treeson2 := &TreeNode{Value: 5, Left:  nil, Right: nil,}
	Tree := &TreeNode{Value: -10, Left:  Treeson1, Right: Treeson2,}
	sum := TreeNodeSumRoot(Tree)
	fmt.Println(sum)
}