package _144_binary_tree_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func CreateTree(treeArr []int) *TreeNode {
	if len(treeArr) == 0 {
		return nil
	}
	root := &TreeNode{Val: treeArr[0]}
	queue := []*TreeNode{root}
	i := 1
	for i < len(treeArr) {
		node := queue[0]
		queue = queue[1:]
		if treeArr[i] != -1 {
			node.Left = &TreeNode{Val: treeArr[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(treeArr) && treeArr[i] != -1 {
			node.Right = &TreeNode{Val: treeArr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}
