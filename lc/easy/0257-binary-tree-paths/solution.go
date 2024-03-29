package _257_binary_tree_paths

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	dfsAns := make([][]int, 0)
	cur := make([]int, 0)
	dfs(&dfsAns, &cur, root)
	return write(dfsAns)
}

func dfs(res *[][]int, cur *[]int, node *TreeNode) {
	if node == nil {
		return
	}
	*cur = append(*cur, node.Val)
	if isLeaf(node) {
		curAns := make([]int, len(*cur))
		copy(curAns, *cur)
		*res = append(*res, curAns)
	} else {
		dfs(res, cur, node.Left)
		dfs(res, cur, node.Right)
	}
	*cur = (*cur)[:len(*cur)-1]
}

func write(dfsAns [][]int) []string {
	res := make([]string, 0)
	for _, item := range dfsAns {
		str := ""
		for i := 0; i < len(item)-1; i++ {
			str += strconv.Itoa(item[i]) + "->"
		}
		str += strconv.Itoa(item[len(item)-1])
		res = append(res, str)
	}
	return res
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}
