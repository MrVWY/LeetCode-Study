package main

import "fmt"

//二叉树路径和
type treeNode struct {
	Value int
	Left *treeNode
	Right *treeNode
}

func Start04(root *treeNode, sum int) [][]int {
	var r [][]int
	if root == nil{
		return r
	}
	var temp []int
	return Walk(root, sum, temp)
}

func Walk(root *treeNode, sum int, temp []int) [][]int {
	var r [][]int
	if root == nil{
		return r
	}
	temps := make([]int,len(temp))
	copy(temps,temp)
	temps = append(temp,root.Value) // 将该根结点的值加到temps
	if root.Left == nil && root.Right == nil{
		if root.Value == sum{
			r = append(r , temps)  //如果遍历到一个节点等于sum ，那么将该路径temps加入到r中
			return r
		}
		return r
	}

	TL := Walk(root.Left, sum - root.Value, temps)  //遍历左子树
	TR := Walk(root.Left, sum - root.Value, temps)  //遍历右子树

	if len(TL) > 0 {
		r = append(r,TL...)
	}

	if len(TR) > 0 {
		r = append(r,TR...)
	}
	return r
}

func main() {
	Treeson1_2 := &treeNode{Value: 7, Left:  nil, Right: nil,}
	Treeson1_1 := &treeNode{Value: 15, Left:  nil, Right: nil,}
	Treeson1 := &treeNode{Value: 20, Left:  Treeson1_1, Right: Treeson1_2,}
	Treeson2 := &treeNode{Value: 5, Left:  nil, Right: nil,}
	Tree := &treeNode{Value: 10, Left:  Treeson1, Right: Treeson2,}
	a := Start04(Tree,45)
	fmt.Println(a)
}