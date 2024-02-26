package _102_binary_tree_level_order_traversal

import "container/list"

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	// for init
	res := make([][]int, 0)
	curLevelCount := 1
	nextLevelCount := 0
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		curLevelRes := make([]int, 0)
		for i := 0; i < curLevelCount; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			curLevelRes = append(curLevelRes, node.Val)
			if node.Left != nil {
				nextLevelCount++
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				nextLevelCount++
				queue.PushBack(node.Right)
			}
		}
		res = append(res, curLevelRes)
		curLevelCount = nextLevelCount
		nextLevelCount = 0
	}
	return res
}
