package _235_lowest_common_ancestor_of_a_binary_search_tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if p == nil || q == nil || root == nil {
		return nil
	}
	if p.Val > q.Val {
		p, q = q, p
	}
	cur := root
	for cur != nil {
		if cur.Val == p.Val || cur.Val == q.Val {
			return cur
		}
		if cur.Val > p.Val && q.Val > cur.Val {
			return cur
		}
		if p.Val > cur.Val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	return cur
}
