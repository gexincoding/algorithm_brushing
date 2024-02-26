package _094_binary_tree_inorder_traversal

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	root := NewNode(1)
	root.Right = NewNode(2)
	root.Right.Left = NewNode(3)
	res := inorderTraversal2(root)
	for _, cur := range res {
		fmt.Print(cur, " ")
	}

}
