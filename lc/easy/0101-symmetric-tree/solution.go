package _101_symmetric_tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root, root)
}

func isMirror(tree1 *TreeNode, tree2 *TreeNode) bool {
	if (tree1 == nil) != (tree2 == nil) {
		return false
	}
	if tree1 == nil && tree2 == nil {
		return true
	}
	if tree1.Val != tree2.Val {
		return false
	}
	return isMirror(tree1.Left, tree2.Right) && isMirror(tree1.Right, tree2.Left)
}
