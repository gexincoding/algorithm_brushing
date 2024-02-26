package _094_binary_tree_inorder_traversal

import "container/list"

func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	stack := list.New()
	cur := root
	for !(cur == nil && stack.Len() == 0) {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}
