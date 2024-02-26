package _110_balanced_binary_tree

type Answer struct {
	depth      int
	isBalanced bool
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root).isBalanced
}

func dfs(tree *TreeNode) Answer {
	if tree == nil {
		return Answer{
			depth:      0,
			isBalanced: true,
		}
	}
	if isLeaf(tree) {
		return Answer{
			depth:      1,
			isBalanced: true,
		}
	}
	res := Answer{}
	ansL := dfs(tree.Left)
	ansR := dfs(tree.Right)
	if ansR.isBalanced && ansL.isBalanced {
		if absInt(ansL.depth-ansR.depth) > 1 {
			res.isBalanced = false
		} else {
			res.isBalanced = true
		}
	} else {
		res.isBalanced = false
	}

	res.depth = getMax(ansL.depth, ansR.depth) + 1
	return res
}

func isLeaf(tree *TreeNode) bool {
	return tree != nil && tree.Left == nil && tree.Right == nil

}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
