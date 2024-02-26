package _583_kth_largest_sum_in_a_binary_tree

import "testing"

func TestSolution(t *testing.T) {
	tree := CreateTree([]int{5, 8, 9, 2, 1, 3, 7, 4, 6})
	kthLargestLevelSum(tree, 2)
}
