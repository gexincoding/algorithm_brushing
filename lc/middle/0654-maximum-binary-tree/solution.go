package _654_maximum_binary_tree

import (
	"math"
)

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return dfs(0, len(nums)-1, nums)
}

func dfs(l, r int, arr []int) *TreeNode {
	if l < r {
		return nil
	}
	index, maxVal := getMax(l, r, arr)
	root := newNode(maxVal)
	root.Left = dfs(l, index-1, arr)
	root.Right = dfs(index+1, r, arr)
	return root

}

func getMax(l, r int, arr []int) (int, int) {
	max := math.MinInt64
	index := -1
	for i := l; i <= r; i++ {
		if arr[i] > max {
			max = arr[i]
			index = i
		}
	}
	return index, max
}

func newNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}
