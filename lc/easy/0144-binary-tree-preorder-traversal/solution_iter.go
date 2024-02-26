package _144_binary_tree_preorder_traversal

import "container/list"

func preorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() != 0 {
		cur := stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, cur.Val)
		if cur.Right != nil {
			stack.PushBack(cur.Right)
		}
		if cur.Left != nil {
			stack.PushBack(cur.Left)
		}
	}
	return res
}
