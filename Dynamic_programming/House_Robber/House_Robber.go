package House_Robber

func max(a,b int) int {
	if a > b {
		return a
	}
	return b

}

//198. 打家劫舍
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int,len(nums)+1)
	dp[1] = nums[0]
	for i := 2 ; i <= len(nums) ; i++ {
		dp[i] = max(dp[i-1],dp[i-2]+nums[i-1]) //注意dp数组比nums数组多一格 dp[0]表示第0间房子可以偷窃0金币
	}
	return dp[len(nums)]
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int,len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0],nums[1])
	for i := 2 ; i < len(nums) ; i++ {
		dp[i] = max(dp[i-1],dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

//213. 打家劫舍 II
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 0 {
		return 0
	}
	return max(robs(nums, 0, len(nums)-1),robs(nums, 1, len(nums)))
}

func robs(nums []int, start, end int) int {
	dp_i, dp_i_1, dp_i_2 := 0, 0, 0
	//    dp_i,   dp_i_1,  dp_i_2
	for i := start ; i < end ; i++ {
		dp_i = max(dp_i_1, nums[i]+dp_i_2)
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i
}

//337. 打家劫舍 III
func rob(root *TreeNode) int {
	selects, unselects := d(root)
	return max(selects, unselects)
}

func d(root *TreeNode) (selects, unselects int) {
	if root == nil {
		return 0,0
	}
	selectsL, unselectsL := d(root.Left)
	selectsR, unselectsR := d(root.Right)

	selects = root.Val + unselectsL + unselectsR
	unselects = max(selectsL,unselectsL) + max(selectsR, unselectsR)
	return selects, unselects
}