package main

import "fmt"

//low[u]:u或u的子树能够追溯到的最早的栈中节点的次序号。
//dfn[u]:节点u搜索的次序编号(时间戳)
//slack栈:回溯时可以判断栈顶到栈中的节点是否为一个强连通分量。
//当DFN(u)=Low(u)时，以u为根的搜索子树上所有节点是一个强连通分量。
var low, dfn [100]int
var slack []int
var componentNumber int //有向图强连通分量个数
var Index int //‘时间戳’
//var E [][]int //图
var visited [100]bool
var comPone [10][]int //结果集

var E = [][]int{{}, {2}, {3}, {4}, {5}, {6}, {7},{1},{}}

func tarJin(node int) {
	low[node] = Index+1
	dfn[node] = Index+1
	visited[node] = true
	slack = append(slack, node)

	for i := 0 ; i < len(E[node]) ; i++ {
		next := E[node][i]
		if dfn[node] == 0 {
			tarJin(next) //下一个节点
			low[i] = min(low[i], low[next])
		}else if visited[node] {
			low[i] = min(low[i], dfn[next])
		}
	}
	//回溯
	if dfn[node] == low[node] {
		componentNumber++
		for {//do
			now := slack[len(slack)-1] //将now退栈，为该强连通分量中一个顶点
			fmt.Print(now)
			slack = slack[:len(slack)-1]
			visited[node] = false
			//while
			if now == node {
				break
			}
		}
	}

}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	//6个顶点, 8条边
	for i := 1 ; i <= 8 ; i++ {
		if dfn[i] == 0 {
			tarJin(i)
		}
	}
}