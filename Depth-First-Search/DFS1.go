package Depth_First_Search

//979. 在二叉树中分配硬币
//定义dfs(node) 为这个节点所在的子树中金币的 过载量,也就是这个子树中金币的数量减去这个子树中节点的数量
//如果树的叶子仅包含 0 枚金币（与它所需相比，它的 过载量 为 -1），那么我们需要从它的父亲节点移动一枚金币到这个叶子节点上  (移动次数+1)
//如果说，一个叶子节点包含 4 枚金币（它的 过载量 为 3），那么我们需要将这个叶子节点中的 3 枚金币移动到别的地方去    (移动次数+3)
func distributeCoins(root *TreeNode) int {
	var count int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		L, R := dfs(node.Left), dfs(node.Right)
		count += abs(L) + abs(R)
		return node.Val + L + R -1
	}
	dfs(root)
	return count
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
