package main

//03. 数组中重复的数字
func findRepeatNumber(nums []int) int {
	m := make(map[int]int, 0)
	res := 0
	for i := 0 ; i < len(nums) ; i++ {
		if _, ok := m[nums[i]] ; ok{
			res = i
			break
		}else{
			m[nums[i]]++
		}
	}
	return nums[res]
}

//04. 二维数组中的查找
//从右上角开始
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows, vers := len(matrix), len(matrix[0])
	row, ver := 0, vers-1
	for row < rows && ver >= 0 {
		a := matrix[row][ver]
		if a == target {
			return true
		}else if a > target {
			ver--
		}else {
			row++
		}
	}
	return false
}

//05. 替换空格
func replaceSpace(s string) string {
	result := ""
	for _ , v := range s {
		switch v {
		case ' ':
			result += "%20"
		default:
			result += string(v)
		}
	}
	return result
}

//06. 从尾到头打印链表
//先反转链表 or 最后反转数组
func reversePrint(head *ListNode) []int {
	res := make([]int, 0)
	if head == nil {
		return res
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		mid := cur.Next
		cur.Next = pre
		pre = cur
		cur = mid
	}

	for pre != nil {
		res = append(res, pre.Val)
		pre = pre.Next
	}
	return res
}

//11. 旋转数组的最小数字
//7 0 1 1 1 1 1 2 3 4
func minArray(numbers []int) int {
	i, j := 0, len(numbers)-1
	for i < j {
		mid := (i+j) >> 1
		if numbers[mid] > numbers[j]{
			i = mid + 1
		}else if numbers[mid] < numbers[j]{
			j = mid
		}else{
			j -= 1
		}
	}
	return numbers[i]
}


//54. 二叉搜索树的第k大节点
//中序遍历的顺序为左中右，即得到的是一个递增序列
//逆中序遍历的顺序为右中左，即得到的是一个递减序列
var count int
var res int
func kthLargest(root *TreeNode, k int) int {
	count = k
	res = 0
	dfs(root)
	return res
}

func dfs(node *TreeNode) {
	if node != nil {
		//逆中序遍历的顺序,因为左边是小于根节点
		dfs(node.Right)
		count--
		if count == 0 {
			res = node.Val
			return
		}
		dfs(node.Left)
	}

}
