package _236_lowest_common_ancestor_of_a_binary_tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, _, res := dfs(root, p, q)
	return res
}

func dfs(root, a, b *TreeNode) (bool, bool, *TreeNode) {
	if root == nil {
		return false, false, nil
	}
	LfindA, LfindB, Lans := dfs(root.Left, a, b)
	RfindA, RfindB, Rans := dfs(root.Right, a, b)
	var findA, findB bool
	findA = root == a || LfindA || RfindA
	findB = root == b || LfindB || RfindB

	if Lans != nil {
		return findA, findB, Lans
	}
	if Rans != nil {
		return findA, findB, Rans
	}
	if findA && findB {
		return findA, findB, root
	}
	return findA, findB, nil
}
