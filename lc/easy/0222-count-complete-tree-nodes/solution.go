package _222_count_complete_tree_nodes

import "math"

func countNodes(root *TreeNode) int {
	return dfs(root)
}

func dfs(tree *TreeNode) int {
	res := 0
	if tree == nil {
		return 0
	}
	if isLeaf(tree) {
		return 1
	}
	fullL, numL := isFull(tree.Left)
	fullR, numR := isFull(tree.Right)
	if fullL {
		res += power(2, numL) - 1
	} else {
		res += dfs(tree.Left)
	}
	if fullR {
		res += power(2, numR) - 1
	} else {
		res += dfs(tree.Right)
	}
	return res + 1

}

func isLeaf(tree *TreeNode) bool {
	return tree != nil && tree.Left == nil && tree.Right == nil
}

// 已经是完全了，判断是不是满！
func isFull(tree *TreeNode) (bool, int) {
	if tree == nil {
		return false, -1
	}
	l, r := 0, 0
	cur := tree
	for cur != nil {
		cur = cur.Left
		l++
	}
	cur = tree
	for cur != nil {
		cur = cur.Right
		r++
	}
	return l == r, l
}

func power(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
