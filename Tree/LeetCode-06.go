package main

import (
	"fmt"
)

type treeNode06 struct {
	Value int
	Left *treeNode06
	Right *treeNode06
}

func Start06(n int) []*treeNode06 {
	var t []*treeNode06
	if n < 1 {
		return t
	}
	return BST(1,n)
}

func BST(start, end int) []*treeNode06 {
	var t []*treeNode06
	if end < start {
		return t
	}
	if start == end {
		t = append(t, &treeNode06{Value:start})
		return t
	}
	if start + 1 == end {
		t = append(t, &treeNode06{Value:start,Left:nil,Right:&treeNode06{Value: end, Left:  nil, Right: nil,}})
		t = append(t, &treeNode06{Value:end,Left:&treeNode06{Value: start, Left:  nil, Right: nil,},Right:nil})
		return t
	}
	for i := 0 ; i <= end ; i ++ {
		L := BST(start, i-1)  // start < i < i-1 , 进行遍历 ，返回来的是数组[]*treeNode06
		R := BST(i+1, end)   // i+1 < i <end , 进行遍历 ，返回来的是数组[]*treeNode06
		if len(L) <= 0 {
			for _, v := range L {
				root := &treeNode06{Value: i, Left:v}
				t = append(t, root)
			}
		}else if len(R) <= 0 {
			for _, v := range R {
				root := &treeNode06{Value:i,Right:v}
				t = append(t, root)
			}
		}else {
			for _, L := range L {
				for _, R := range R {
					root := &treeNode06{Value:i,Left:L,Right:R}
					t = append(t, root)
				}
			}
		}
	}
	return t
}

func main() {
	a := Start06(1)
	fmt.Println(a)
}