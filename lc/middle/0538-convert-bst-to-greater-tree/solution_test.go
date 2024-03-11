package _538_convert_bst_to_greater_tree

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	root := CreateTree([]int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8})
	bst := convertBST(root)
	fmt.Println(bst)
}
