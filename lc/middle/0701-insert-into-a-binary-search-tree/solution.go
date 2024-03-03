package _701_insert_into_a_binary_search_tree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return newNode(val)
	}
	cur := root
	var pre *TreeNode
	var op bool
	for cur != nil {
		pre = cur

		if val < cur.Val {
			op = true
			cur = cur.Left
		} else {
			op = false
			cur = cur.Right
		}
	}
	if pre == nil {
		return newNode(val)
	} else {
		if op {
			pre.Left = newNode(val)
		} else {
			pre.Right = newNode(val)
		}
	}

	return root
}

func newNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}
