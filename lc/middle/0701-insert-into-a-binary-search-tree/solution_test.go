package _701_insert_into_a_binary_search_tree

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	tree := CreateTree([]int{4, 2, 7, 1, 3})
	bst := insertIntoBST(tree, 5)
	fmt.Println(bst)
}
