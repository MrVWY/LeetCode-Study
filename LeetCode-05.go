package main

import "fmt"

//验证是否为二叉树
type treeNode05 struct {
	Value int
	Left *treeNode05
	Right *treeNode05
}

func maxBFS(root *treeNode05) int {
	if root.Right != nil{
		return maxBFS(root.Right)
	}
	return root.Value
}

func minBFS(root *treeNode05) int {
	if root.Left != nil {
		return minBFS(root.Left)
	}
	return root.Value
}

func isValiBFS(root *treeNode05) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Right != nil {
		L := maxBFS(root.Left) // 找到左子节点的右节点中最大的
		R := minBFS(root.Right)  // 找到右子节点的左节点中最小的
		return isValiBFS(root.Left) && isValiBFS(root.Right) && L < root.Value && root.Value < R
	}
	if root.Left != nil {
		L := maxBFS(root.Left) // 找到左子节点的右节点中最大的
		return isValiBFS(root.Left) && L < root.Value
	}
	if root.Right != nil {
		R := minBFS(root.Right)
		return isValiBFS(root.Right) && root.Value < R
	}
	return true
}

func main() {
	Treeson1_2 := &treeNode05{Value: 7, Left:  nil, Right: nil,}
	Treeson1_1 := &treeNode05{Value: 15, Left:  nil, Right: nil,}
	Treeson1 := &treeNode05{Value: 20, Left:  Treeson1_1, Right: Treeson1_2,}
	Treeson2 := &treeNode05{Value: 5, Left:  nil, Right: nil,}
	Tree := &treeNode05{Value: 10, Left:  Treeson1, Right: Treeson2,}
	a := isValiBFS(Tree)
	fmt.Println(a)
}