package main

import "fmt"
//输入: [1,2,3,null,5,null,4]
//输出: [1, 3, 4]
//解释:
//
//             1            <---
//           /   \
//          2     3         <---
//           \     \
//            5     4       <---

//解题思路：
//1，这个题目是层序遍历的变形：层序遍历不需要每一层断开，本题需要每一层断开，取每一层最右元素（左视图类似）
//2，这里需要两个队列：
//   A，第一个队列存上一层结果
//   B，第二个队列存上一层的所有孩子
//3，每次取上一层最右元素

type treeNode07 struct {
	Value int
	Left *treeNode07
	Right *treeNode07
}

//定义一个队列
type queue struct {
	data []*treeNode07
}

func (q *queue) last() *treeNode07 {
	L := len(q.data)
	var r *treeNode07
	if L > 0 {
		r = q.data[L - 1]
	}
	return r
}

func (q *queue) push(r *treeNode07) {
	q.data = append(q.data, r)
}

func (q *queue) pop() *treeNode07 {
	L := len(q.data)
	var r  *treeNode07
	if L > 0 {
		r = q.data[0]
		q.data = q.data[1:]
	}
	return r
}

func (q *queue) emtry() bool {
	return len(q.data) <= 0
}

func RightSideView(root *treeNode07) []int {
	var q queue //定义一个队列,主要存放root的这一层的节点
	var rr []int
	if root != nil {
		q.push(root)  //队列q包含了该树的所有节点 q.data = [(1,Left,Right)]
	}
	for !q.emtry() {
		L := q.last() // 取队列q的最后一个节点
		rr = append(rr, L.Value)  //(1)rr = [1]  (2) rr = [1,3] (4) rr = [1,3,4]
		var q1 queue  // 定义一个队列,主要存放root的下一层的节点
		for !q.emtry() {
			r := q.pop() //取出队列q的root节点的下一层节点,此时q.data是[0] ,第二次遍历时q.data为[1]
			//如果r.Left等于nil,那么就会出去r.Right的queue
			if r.Left !=nil {
				q1.push(r.Left)  //r.Left 是一个 *treeNode07
			}
			//如果r.Right等于nil,那么就会出去r.Left的queue
			if r.Right != nil {
				q1.push(r.Right)
			}
			//(1)q1.data(1,Left,Right) = [(2,Left,Right),(3,Left,Right)]
			//(2)q1.data = [(5,Left,Right),(4,Left,Right)]
		}
		q = q1
	}
	return rr
}

func main() {
	Treeson1_2 := &treeNode07{Value: 4, Left:  nil, Right: nil,}
	Treeson1_1 := &treeNode07{Value: 5, Left:  nil, Right: nil,}
	Treeson1 := &treeNode07{Value: 2, Left:  nil, Right: Treeson1_1,}
	Treeson2 := &treeNode07{Value: 3, Left:  nil, Right: Treeson1_2,}
	Tree := &treeNode07{Value: 1, Left:  Treeson1, Right: Treeson2,}
	a := RightSideView(Tree)
	fmt.Println(a)
}