package _583_kth_largest_sum_in_a_binary_tree

import (
	"fmt"
	"testing"
)

func TestCreateTree(t *testing.T) {
	tree := CreateTree([]int{1, 2, 3, 4, -1, 5})
	fmt.Println(tree)
}
