package _530_minimum_absolute_difference_in_bst

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	tree := CreateTree([]int{5, 4, 7})
	fmt.Println(getMinimumDifference(tree))

}
