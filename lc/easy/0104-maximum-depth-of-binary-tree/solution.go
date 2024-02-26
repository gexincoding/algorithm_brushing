package _104_maximum_depth_of_binary_tree

func maxDepth(root *TreeNode) int {
	return maxH(root)
}

func maxH(tree *TreeNode) int {
	if tree == nil {
		return 0
	}
	return getMax(maxH(tree.Right), maxH(tree.Left)) + 1
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
