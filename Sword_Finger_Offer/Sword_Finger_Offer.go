package Sword_Finger_Offer

import "math"

type ListNode struct {
	Val int
	Next *ListNode
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
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

//07. 重建二叉树
//[ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
//[ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := 0
	for k := range inorder {
		if inorder[k] == preorder[0] {
			root = k
			break
		}
	}
	return &TreeNode {
		Val : preorder[0],
		Left : buildTree(preorder[1:root+1], inorder[0:root]),
		Right : buildTree(preorder[root+1:],inorder[root+1:]),
	}
}

//09. 用两个栈实现队列
type CQueue struct {
	Stack1 []int //插入队列
	Stack2 []int //删除队列
}

func Constructor() CQueue {
	return CQueue{
		Stack1 : make([]int, 0),
		Stack2 : make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int)  {
	this.Stack1 = append(this.Stack1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.Stack2) == 0 {
		for len(this.Stack1) > 0 {
			this.Stack2 = append(this.Stack2, this.Stack1[len(this.Stack1)-1])
			this.Stack1 = this.Stack1[:len(this.Stack1)-1]
		}
	}
	if len(this.Stack2) != 0 {
		value := this.Stack2[len(this.Stack2)-1]
		this.Stack2 = this.Stack2[:len(this.Stack2)-1]
		return value
	}
	return -1
}



//11. 旋转数组的最小数字
//二分
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

//27. 二叉树的镜像
//想象只有三个节点的树的镜像
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	tmp := root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(tmp)
	return root
}

//30. 包含min函数的栈
//Mins的第i位对应Slack的前i位的最小值
type MinStack struct {
	Slack []int
	Mins  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack {
		Slack : make([]int,0),
		Mins : []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int)  {
	this.Slack = append(this.Slack, x)
	this.Mins = append(this.Mins, min(x, this.Mins[len(this.Mins)-1]))
}

func (this *MinStack) Pop()  {
	this.Slack = this.Slack[:len(this.Slack)-1]
	this.Mins = this.Mins[:len(this.Mins)-1]
}

func (this *MinStack) Top() int {
	return this.Slack[len(this.Slack)-1]
}

func (this *MinStack) Min() int {
	return this.Mins[len(this.Mins)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//31. 栈的压入、弹出序列
//定义一个辅助栈来模拟压栈和出栈
func validateStackSequences(pushed []int, popped []int) bool {
	slack := make([]int, 0)
	i := 0
	for _, v := range pushed {
		slack = append(slack, v)
		for len(slack) != 0 && slack[len(slack)-1] == popped[i] {
			slack = slack[:len(slack)-1]
			i++
		}
	}

	return len(slack) == 0
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


//60. n个骰子的点数
func twoSum(n int) []float64 {
	dp := make([][]int, n+1)
	for i := 0 ; i < len(dp) ; i++ {
		dp[i] = make([]int, 6*n+1)
	}
	for i := 1 ; i <= 6 ; i++{
		dp[1][i] = 1
	}
	for i := 1 ; i <= n ; i++ {
		dp[i][i] = 1
	}

	for i := 2 ; i <= n ; i++ {//几个骰子
		for j := i+1 ; j <= 6*i ; j++ {//点数
			for k := 1 ; k <= 6 ; k++ {
				if j - k <=0 {break}
				dp[i][j] += dp[i-1][j-k]
			}
		}
	}
	res := make([]float64, 6*n+1)
	for i := n ; i <= 6*n ; i++{
		res[i] = float64(dp[n][i])/ math.Pow(6,float64(n))
	}
	return res[n:]
}

//61. 扑克牌中的顺子
func isStraight(nums []int) bool {
	m := make(map[int]bool)
	max, min := 0, 14
	for i := 0 ; i < len(nums) ; i++ {
		if _, ok := m[nums[i]]; ok {
			return false
		}
		if nums[i] == 0 {
			continue
		}
		m[nums[i]] = true
		min = mins(min, nums[i])
		max = maxs(max, nums[i])
	}
	//fmt.Prinln(max, min)
	return max - min < 5
}

func mins(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxs(a, b int) int {
	if a > b {
		return a
	}
	return b
}