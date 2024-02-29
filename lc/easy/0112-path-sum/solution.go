package _112_path_sum

func hasPathSum(root *TreeNode, targetSum int) bool {

	return dfs(root, targetSum)
}

// 当前节点是node 问 node这个棵树，有没有target的路径
func dfs(tree *TreeNode, tar int) bool {

	if tree == nil {
		return false
	}
	if isLeaf(tree) {
		return tree.Val == tar
	}
	return dfs(tree.Left, tar-tree.Val) || dfs(tree.Right, tar-tree.Val)
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}
