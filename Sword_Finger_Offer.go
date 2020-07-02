package main

type ListNode struct {
	Val int
	Next *ListNode
}

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

//16. 数值的整数次方
//快速幂
//虑比如我们计算 x^8，就是 x^2 * x^2 * x^2 * x^2，当我们计算出来 x^2 之后就可以只进行三次乘法就可以了，相对于之前的 7 次乘法，时间大大减少了。
//也就是 x^n 可以分解成若干个 x^i 的乘积
//我们这里使用快速幂进行求解。我们看一下 n 的二进制形式一定是若干个 1 和 0 构成，比如 9 = 1001 = 1*2^3 + 0*2^2 + 0*2^1 + 1*2^0
//所以我们可以看出来，每次乘的值都是前一个值的2倍，当 n 对应位为0时跳过
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	res := 1.00
	if n < 0 {
		x, n = 1/x, -n
	}

	for n != 0 {
		//判断该位是否为1
		if n & 1 == 1 {
			res = res * x
		}
		x = x * x
		n = n >> 1
	}
	return res
}

//18. 删除链表的节点
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Val:0 , Next: head}
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}
	return dummy.Next
}

//21. 调整数组顺序使奇数位于偶数前面
func exchange(nums []int) []int {
	left, right :=0, len(nums)-1
	for left < right {
		if nums[left] & 1 == 1 {
			left++
			continue
		}

		if nums[right] & 1 == 0 {
			right--
			continue
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}

//22. 链表中倒数第k个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for i := 0 ; i < k ; i++ {
		if fast == nil {
			return slow
		}
		fast = fast.Next
	}

	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}

//25. 合并两个排序的链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummy := &ListNode{} //虚拟节点
	head := dummy //移动节点
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next =  l1
			l1 = l1.Next
		}else  {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}
	return dummy.Next
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
