package _094_binary_tree_inorder_traversal

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	dfs(root, &res)
	return res
}

func dfs(cur *TreeNode, res *[]int) {
	if cur == nil {
		return
	}
	dfs(cur.Left, res)
	*res = append(*res, cur.Val)
	dfs(cur.Right, res)
}
