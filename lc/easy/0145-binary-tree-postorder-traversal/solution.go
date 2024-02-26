package _145_binary_tree_postorder_traversal

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	dfs145(root, &res)
	return res
}

func dfs145(cur *TreeNode, res *[]int) {
	if cur == nil {
		return
	}
	dfs145(cur.Left, res)
	dfs145(cur.Right, res)
	*res = append(*res, cur.Val)
}
