package Daily

import "sort"

//1608. 特殊数组的特征值
func specialArray(nums []int) int {
	sort.Ints(nums)
	count := 0
	var flag bool
	for i := 0 ; i <= nums[len(nums)-1] ; i++ {
		count = 0
		for j := 0 ; j < len(nums) ; j++ {
			if nums[j] >= i {
				count++
			}
		}
		if count == i {
			flag = true
			break
		}
	}
	if flag {
		return count
	}
	return -1
}

//1609. 奇偶树
/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}
func isEvenOddTree(root *TreeNode) bool {
	odd := make([][]int, 0)
	even := make([][]int, 0)
	level := 1
	queue := make([]*TreeNode,0)
	queue = append(queue, root)
	for len(queue) > 0 {
		tmp := queue
		queue = nil
		o := make([]int, 0)
		e := make([]int, 0)
		for len(tmp) > 0 {
			v := tmp[0]
			tmp = tmp[1:]
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
			if level % 2 == 0 {
				if v.Val % 2 == 0 {
					o = append(o, v.Val)
				}else{
					return false
				}
			}else {
				if v.Val % 2 != 0 {
					e = append(e, v.Val)
				}else{
					return false
				}
			}
		}
		odd = append(odd, o)
		even = append(even, e)
		level++
	}
	//fmt.Println(odd, even)
	for _, v := range odd {
		for i := 0 ; i < len(v)-1 ; i++ {
			if v[i+1] >= v[i] {
				return false
			}
		}
	}
	for _, v := range even {
		for i := 0 ; i < len(v)-1 ; i++ {
			if v[i+1] <= v[i] {
				return false
			}
		}
	}
	return true
}