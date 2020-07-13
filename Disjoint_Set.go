package main

import "fmt"

func fundRoot(node int, parent []int) int {
	xRoot := node
	for parent[xRoot] != -1 {
		xRoot = parent[xRoot]
	}
	return xRoot
}


func UnionVertices(node1, node2 int, parent, rank []int) bool {
	xRoot := fundRoot(node1, parent)
	yRoot := fundRoot(node2, parent)
	if xRoot == yRoot {
		return false //union is fail
	}else {
		if rank[xRoot] > rank[yRoot]  {  // 让 少的指向多 的
			parent[yRoot] = xRoot
		} else if rank[xRoot] < rank[yRoot] {
			parent[xRoot] = yRoot
		} else {
			parent[xRoot] = yRoot   // 这个随便可以
			rank[yRoot]++
		}
		return true
	}
}

func main() {
	parent := make([]int, 6)
	rank := make([]int, 6) //深度
	for i := 0 ; i < 6 ; i++ {
		parent[i] = -1
	}
	//判断有没有环路
	edges := [][]int{
		{0,1},{1,2},{1,3},{2,4},{3,4},{2,5},
	}
	for i := 0; i < 6; i++ {
		x := edges[i][0]
		y := edges[i][1]
		if UnionVertices(x, y, parent,rank) {
			fmt.Println("Cycle detected! ")
			break
		}
	}
}