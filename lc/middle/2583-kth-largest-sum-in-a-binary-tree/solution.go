package _583_kth_largest_sum_in_a_binary_tree

import (
	"container/list"
	"sort"
)

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	if root == nil {
		return -1
	}
	levelSum := make([]int, 0)
	queue := list.New()
	curLevelCount := 1
	curLevelSum := 0
	nextLevelCount := 0
	queue.PushBack(root)
	for queue.Len() != 0 {

		for i := 0; i < curLevelCount; i++ {
			curNode := queue.Remove(queue.Front()).(*TreeNode)
			curLevelSum += curNode.Val
			if curNode.Right != nil {
				nextLevelCount++
				queue.PushBack(curNode.Right)
			}
			if curNode.Left != nil {
				nextLevelCount++
				queue.PushBack(curNode.Left)
			}
		}
		levelSum = append(levelSum, curLevelSum)
		curLevelCount = nextLevelCount
		nextLevelCount = 0
		curLevelSum = 0
	}
	sort.Slice(levelSum, func(i, j int) bool {
		return levelSum[i] > levelSum[j]
	})
	if len(levelSum) < k {
		return -1
	}
	return int64(levelSum[k-1])
}
