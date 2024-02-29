package _513_find_bottom_left_tree_value

import "container/list"

func findBottomLeftValue(root *TreeNode) int {
	// root != nil

	queue := list.New()
	curLevelCount := 1
	nextLevelCount := 0
	curLevelFirstNum := root.Val
	queue.PushBack(root)

	for queue.Len() != 0 {
		for i := 0; i < curLevelCount; i++ {
			curNode := queue.Remove(queue.Front()).(*TreeNode)
			if i == 0 {
				curLevelFirstNum = curNode.Val
			}
			if curNode.Left != nil {
				queue.PushBack(curNode.Left)
				nextLevelCount++
			}
			if curNode.Right != nil {
				queue.PushBack(curNode.Right)
				nextLevelCount++
			}
		}
		curLevelCount = nextLevelCount
		nextLevelCount = 0
	}
	return curLevelFirstNum
}
