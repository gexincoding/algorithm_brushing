package _144_binary_tree_preorder_traversal

import (
	"fmt"
	"testing"
)

func Test114(t *testing.T) {
	root := NewNode(1)
	root.Right = NewNode(2)
	root.Right.Left = NewNode(3)
	res := preorderTraversal2(root)
	for _, cur := range res {
		fmt.Print(cur, " ")
	}

}
