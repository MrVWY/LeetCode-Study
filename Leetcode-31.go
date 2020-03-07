package main

import (
	"bytes"
	"fmt"
	"sort"
)

//1029. 两地调度  思想：贪心,先把所有人2N派去A,在从priceA-priceB中挑最小的N个人去B
type IntSlice []int
func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
func twoCitySchedCost(costs [][]int) int {
	res,cons :=0,[]int{}
	for _, v := range costs {
		res += v[0]
		cons = append(cons, v[1]-v[0])
	}
	sort.Ints(IntSlice(cons))
	fmt.Println(res,cons)
	for i:= 0;i<len(cons)/2;i++ {
		res +=cons[i]
	}

	return res
}

type stack struct {
	item rune
	time int
}

func removeDuplicates(s string, k int) string {
	length := len(s)
	stacks := make([]stack,0,length)
	for _ , v := range s {
		if len(stacks) == 0 || stacks[len(stacks)-1].item != v {
			stacks = append(stacks,stack{item:v,time:1})
		}else {
			stacks[len(stacks)-1].time++
		}

		if stacks[len(stacks)-1].time == k {
			stacks = stacks[:len(stacks)-1]
		}

	}
	runes := make([]rune,0,len(s))
	for _ , item := range stacks {
		for item.time > 0 {
			runes = append(runes,item.item)
			item.time--
		}
	}
	return string(runes)
}

//515. 在每个树行中找最大值
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//以[1,3,2]  分层！！
//-1 2 {root,nil}
//-1 4 {root,   nil,                 root.Left,root.Right}
//     { 1    层与层的分隔               2           2    }
//1	   {root,nil,root.Left,root.Right,nil}
//1 5
//1 5
//1 5
//4	   {root,nil,root.Left,root.Right,nil,nil}
func largestValues(root *TreeNode) []int {
	ret := []int{}
	if root == nil{
		return ret
	}
	array := []*TreeNode{root,nil}
	last := -1
	for {
		if array[last+1] == nil {
			break
		}
		max := array[last+1].Val
		for i := last+1 ; i < len(array) ; i++ {
			fmt.Println(last,len(array)) //一个fmt提交上去用时+2ms................
			if array[i] == nil { //root的左右节点先加进array,在遇到第i层的分隔符，当第i+1层的节点都加进来以后再加上第i+1层的分隔符
			    array = append(array, nil)
				last = i //换层也就是换index
				fmt.Println(last)
				break
			}
			if array[i].Left != nil{
				array = append(array, array[i].Left)
			}
			if array[i].Right != nil{
				array = append(array, array[i].Right)
			}
			if array[i].Val > max{
				max = array[i].Val
			}
		}
		ret = append(ret, max)
	}
	return ret
}

//以[1,3,2]  分层！！
//result = [第一层，第二层 ]
func largestValues(root *TreeNode) []int {
	var (
		result []int
		dfs func(node *TreeNode,depath int)
	)
	dfs = func(node *TreeNode,depath int) {
		if node == nil {
			return
		}

		if depath > len(result) {
			result = append(result,node.Val)
		}else if node.Val > result[depath-1] {
			result[depath-1] = node.Val
		}

		dfs(node.Left,depath+1)
		dfs(node.Right,depath+1)
	}
	return result
}


//706. 设计哈希映射
type MyHashMap struct {
	content map[int]int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
	my := MyHashMap{}
	my.content = map[int]int{}
	return my
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
	this.content[key] = value
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	val, ok := this.content[key]
	if !ok {
		return -1
	}

	return val
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
	delete(this.content, key)
}