package _145_binary_tree_postorder_traversal

import "container/list"

func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() != 0 {
		cur := stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, cur.Val)
		if cur.Left != nil {
			stack.PushBack(cur.Left)
		}
		if cur.Right != nil {
			stack.PushBack(cur.Right)
		}
	}
	res1 := make([]int, 0)
	for i := len(res) - 1; i >= 0; i-- {
		res1 = append(res1, res[i])
	}
	return res1
}
