package _108_convert_sorted_array_to_binary_search_tree

func sortedArrayToBST(nums []int) *TreeNode {
	return dfs(&nums, 0, len(nums)-1)
}
func dfs(nums *[]int, l, r int) *TreeNode {
	if l < r {
		return nil
	}
	mid := (l + r) / 2
	root := newNode((*nums)[mid])
	root.Left = dfs(nums, l, mid-1)
	root.Right = dfs(nums, mid+1, r)
	return root

}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
