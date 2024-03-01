package _106_construct_binary_tree_from_inorder_and_postorder_traversal

func buildTree(in []int, post []int) *TreeNode {
	if post == nil || len(post) == 0 {
		return nil
	}
	if len(post) == 1 {
		return &TreeNode{
			Val: post[0],
		}
	}
	root := &TreeNode{
		Val: post[len(post)-1],
	}
	var lin, lpost, rin, rpost []int
	for i := 0; i < len(in); i++ {
		if in[i] == post[len(post)-1] {
			lin, lpost = in[0:i], post[0:i]
			rin, rpost = in[i+1:], post[i:len(post)-1]
			break
		}
	}
	root.Left, root.Right = buildTree(lin, lpost), buildTree(rin, rpost)
	return root
}
