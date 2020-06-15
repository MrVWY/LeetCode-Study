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

//后继者  二叉搜索树
//所谓 p 的后继节点，就是这串升序数字中，比 p 大的下一个。
//如果 p 大于当前节点的值，说明后继节点一定在 RightTree
//如果 p 等于当前节点的值，说明后继节点一定在 RightTree
//如果 p 小于当前节点的值，说明后继节点一定在 LeftTree 或自己就是
		//递归调用 LeftTree，如果是空的，说明当前节点就是答案
		//如果不是空的，则说明在 LeftTree 已经找到合适的答案，直接返回即可
//二叉搜索树 中序遍历  left mid right
//右子树的左子树前面永远的mid
//左子树的右子树后面永远是mid
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val <= p.Val {
		return inorderSuccessor(root.Right, p)
	}else{
		left := inorderSuccessor(root.Left, p)
		//重点
		if left == nil {
			return root
		}else {
			return left
		}
	}

	return nil
}
