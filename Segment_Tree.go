package main

import "fmt"

func buildTree(tree, arr []int, node, start, end int) {
	if start == end {
		tree[node] = arr[start]
	}else {
		mid := (start+end) >> 1
		leftNode := node * 2 + 1
		rightNode := node * 2 + 2
		buildTree(tree, arr, leftNode, start, mid)
		buildTree(tree, arr, rightNode, mid+1, end)
		tree[node] = tree[leftNode] + tree[rightNode]
	}
}

func updateTree(tree, arr []int, node, start, end, idx, val int) {
	if start == end {
		arr[idx] = val
		tree[node] = val
	}else {
		mid := (start+end) >> 1
		leftNode := node * 2 + 1
		rightNode := node * 2 + 2
		if idx >= start && idx <= mid {
			updateTree(tree, arr, leftNode, start, mid, idx, val)
		}else {
			updateTree(tree, arr, rightNode, mid+1, end, idx, val)
		}
		tree[node] = tree[leftNode] + tree[rightNode]
	}
}

func queryTree(tree, arr []int, node, start, end, L, R int) int{
	if start > R || end < L {
		return 0
	}else if L <= start && R <= end{
		//start,end = 2,5  L,R = 1,5
		return tree[node]
	}else if start == end{
		return tree[node]
	}

	mid := (start+end) >> 1
	leftNode := node * 2 + 1  // node << 1 | 1
	rightNode := node * 2 + 2

	return queryTree(tree, arr, leftNode, start, mid, L, R) + queryTree(tree, arr, rightNode, mid+1, end, L, R)
}

func main() {
	tree := make([]int, 15)
	arr := []int{1,3,5,7,9,11}
	buildTree(tree,arr,0, 0, 5)
	fmt.Println(tree)
	updateTree(tree,arr,0,0,5,4,6)
	fmt.Println(tree)
	fmt.Println(queryTree(tree, arr, 0, 0, 5, 2, 5))
}