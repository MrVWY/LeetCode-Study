package Daily

import (
	"math"
)

//5465. 子树中标签相同的节点数
var ans []int
var graph [][]int
func countSubTrees(n int, edges [][]int, labels string) []int {
	graph = make([][]int, n)
	for i := 0 ; i < n ; i++ {
		graph[i] = make([]int, 0)
	}
	for _, v := range edges {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	ans = make([]int, n)
	dfs(0, make([]bool, n), labels)
	return ans
}

func dfs(node int, visited []bool, labels string) []int {
	visited[node] = true
	count := make([]int, 26)
	count[labels[node]-'a']++
	for _, next := range graph[node] {
		if !visited[next]{
			res := dfs(next, visited, labels)
			for i := 0; i < 26; i++ {
				count[i] += res[i];
			}
		}
	}
	ans[node] = count[labels[node]-'a']
	return count
}


//5467. 找到最接近目标值的函数值 超时
func buildTree(tree, arr []int, node, start, end int) {
	if start == end {
		tree[node] = arr[start]
	}else {
		mid := (start+end) >> 1
		leftNode := node << 1
		rightNode := node << 1 | 1
		buildTree(tree, arr, leftNode, start, mid)
		buildTree(tree, arr, rightNode, mid+1, end)
		tree[node] = tree[leftNode] & tree[rightNode]
	}
}

func queryTree(tree []int, node, start, end, L, R int) int {
	if L == start && R == end {
		return tree[node]
	}
	mid := (start+end) >> 1
	leftNode := node << 1
 	rightNode := node <<1 | 1
 	if R <= mid {
 		return queryTree(tree, leftNode, start, mid, L, R)
	}
	if L > mid {
		return queryTree(tree, rightNode, mid+1, end, L, R)
	}
	return queryTree(tree, leftNode, start, mid, L, mid) & queryTree(tree, rightNode, mid+1, end, mid+1, R)
}

func closestToTarget(arr []int, target int) int {
	n := len(arr)
	tree := make([]int, 100000)
	buildTree(tree, arr, 1, 0, n-1)
	l, r, ans := 0, 0, math.MaxInt32
	for r < n {
		res := queryTree(tree, 1, 0, n-1, l, r)
		if res < target {
			r = max(l+1, r)
		}else {
			r++
		}
		ans = min(ans, int(math.Abs(float64(res-target))))
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {

}