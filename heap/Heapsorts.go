package main

import "fmt"

func heapSort1(tree []int, size int)  {
	lastNode := size - 1
	parent := ( lastNode - 1 ) / 2
	for i := parent ; i >= 0 ; i-- {
		heaping1(tree,i,size)
	}
}

func heaping1(tree []int, root int, size int)  {
	if root > size {
		return
	}
	left := 2*root + 1
	right := 2*root + 2
	//max := tree[root]
	max := root //假设root是最大值
	if left < size && tree[left] > tree[max] {
		fmt.Println(tree[root])
		max = left
	}
	if right < size && tree[right] > tree[max]  {
		fmt.Println(tree[root])
		max = right
	}
	if max != root {
		swap(tree,max,root)
		heaping1(tree,max,size)

	}
}

func swap(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	tree := []int{1,2,3,5,10,4}
	heapSort1(tree, len(tree))
	fmt.Println(tree)
}
