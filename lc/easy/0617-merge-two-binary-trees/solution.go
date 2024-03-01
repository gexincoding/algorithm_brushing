package _617_merge_two_binary_trees

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 != nil && root2 == nil {
		return root1
	}
	if root2 != nil && root1 == nil {
		return root2
	}
	tree := newNode(root1.Val + root2.Val)
	tree.Left = mergeTrees(root1.Left, root2.Left)
	tree.Right = mergeTrees(root1.Right, root2.Right)
	return tree
}

func newNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}
