package _236_lowest_common_ancestor_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
