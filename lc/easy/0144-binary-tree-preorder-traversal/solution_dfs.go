package _144_binary_tree_preorder_traversal

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	dfs(root, &res)
	return res
}

func dfs(cur *TreeNode, res *[]int) {
	if cur == nil {
		return
	}
	*res = append(*res, cur.Val)
	dfs(cur.Left, res)
	dfs(cur.Right, res)
}
