package _098_validate_binary_search_tree

import "math"

func isValidBST(root *TreeNode) bool {
	_, _, isBST := dfs(root)
	return isBST
}

func dfs(tree *TreeNode) (max int, min int, isBST bool) {
	if tree == nil {
		return math.MinInt64, math.MaxInt64, true
	}
	if isLeaf(tree) {
		return tree.Val, tree.Val, true
	}
	lMax, lMin, l := dfs(tree.Left)
	rMax, rMin, r := dfs(tree.Right)
	return getMax(lMax, rMax, tree.Val), getMin(lMin, rMin, tree.Val), lMax < tree.Val && rMin > tree.Val && l && r
}

func isLeaf(tree *TreeNode) bool {
	return tree != nil && tree.Left == nil && tree.Right == nil
}

func getMax(a, b, c int) int {
	if a > b {
		if c > a {
			return c
		} else {
			return a
		}
	} else {
		if c > b {
			return c
		} else {
			return b
		}
	}

}

func getMin(a, b, c int) int {
	if a < b {
		if c < a {
			return c
		} else {
			return a
		}
	} else {
		if c < b {
			return c
		} else {
			return b
		}
	}
}
