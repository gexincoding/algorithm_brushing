package _404_sum_of_left_leaves

func sumOfLeftLeaves(root *TreeNode) int {

	sum := 0
	dfs(&sum, root)
	return sum
}

func dfs(sum *int, node *TreeNode) {
	if node == nil {
		return
	}
	if isLeaf(node) {
		return
	}
	if isLeaf(node.Left) {
		*sum += node.Left.Val
		dfs(sum, node.Right)
	} else {
		dfs(sum, node.Left)
		dfs(sum, node.Right)
	}
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}
