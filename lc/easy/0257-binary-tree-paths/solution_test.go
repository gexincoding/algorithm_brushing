package _257_binary_tree_paths

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {

	tree := CreateTree([]int{1, 2, 3, -1, 5})
	paths := binaryTreePaths(tree)
	fmt.Println(paths)

}
