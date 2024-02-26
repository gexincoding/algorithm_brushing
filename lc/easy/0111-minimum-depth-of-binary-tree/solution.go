package _111_minimum_depth_of_binary_tree

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minR := minDepth(root.Right)
	minL := minDepth(root.Left)
	if minR == 0 {
		return minL + 1
	}
	if minL == 0 {
		return minR + 1
	}
	return getMin(minR, minL) + 1
}

func getMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}
