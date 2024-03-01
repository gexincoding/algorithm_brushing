package _106_construct_binary_tree_from_inorder_and_postorder_traversal

import (
	"fmt"
	"testing"
)

func TestAA(t *testing.T) {
	node := buildTree([]int{1, 2}, []int{1, 2})
	fmt.Println(node)
}
