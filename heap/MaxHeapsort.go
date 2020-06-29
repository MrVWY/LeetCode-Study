package main

import "fmt"

func heapSort(tree []int, size int)  {
	lastNode := size - 1
	parent := ( lastNode - 1 ) / 2
	for i := parent ; i >= 0 ; i-- {
		heaping(tree,i,size)
	}
}

func heaping(tree []int, root int, size int)  {
	if root > size {
		return
	}
	left := 2*root + 1 //左孩子
	right := 2*root + 2 //右孩子
	max := tree[root] //假设root是最大值
	//比较root left right
	if left < size && tree[left] > max {
		max = tree[left]
	}
	if right < size && tree[right] > max  {
		max = tree[right]
	}

	if max != tree[root] {
		if right < size && max == tree[right]  {
			tree[right] = tree[root]
			heaping(tree,right,size)
		}
		if left < size && max == tree[left]  {
			tree[left] = tree[root]
			heaping(tree,left,size)
		}
		tree[root] = max
	}
}

func main() {
	tree := []int{1,2,3,5,10,4}
	heapSort(tree, len(tree))
	fmt.Println(tree)
}