package _530_minimum_absolute_difference_in_bst

import "math"

var pre *TreeNode

func getMinimumDifference(root *TreeNode) int {
	res := math.MaxInt
	pre = nil
	dfs(root, &res)
	return res
}

func dfs(cur *TreeNode, res *int) {
	if cur == nil {
		return
	}
	dfs(cur.Left, res)
	if pre != nil {
		*res = getMin(*res, getAbs(cur.Val-pre.Val))
	}
	pre = cur
	dfs(cur.Right, res)

}

func isLeaf(tree *TreeNode) bool {
	return tree != nil && tree.Right == nil && tree.Left == nil
}

func getMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func getAbs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
