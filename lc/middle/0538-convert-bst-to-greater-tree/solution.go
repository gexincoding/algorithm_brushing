package _538_convert_bst_to_greater_tree

import "log"

var pre = 0

func convertBST(root *TreeNode) *TreeNode {
	pre = 0
	return dfs(root)
}

func dfs(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	r := dfs(root.Right)

	log.Println("pre is: ", pre)
	log.Println("root.val is: ", root.Val)
	root.Val += pre
	root.Right = r
	pre = root.Val

	l := dfs(root.Left)
	root.Left = l

	return root
}
func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
