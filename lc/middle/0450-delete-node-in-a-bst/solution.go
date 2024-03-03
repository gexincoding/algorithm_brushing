package _450_delete_node_in_a_bst

// 看作一个递归函数
// 含义：当前来到的节点是root 需要删除key，返回删除后的BST的头节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if isLeaf(root) {
			return nil
		}
		if leftNil(root) {
			return root.Right
		}
		if rightNil(root) {
			return root.Left
		}
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		cur.Left = root.Left
		return root.Right
	}
	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)
	return root
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}

// 左空右不空
func leftNil(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right != nil
}

// 右空左不空
func rightNil(node *TreeNode) bool {
	return node != nil && node.Left != nil && node.Right == nil
}
