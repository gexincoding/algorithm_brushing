package _501_find_mode_in_binary_search_tree

import (
	"fmt"
	"testing"
)

func TestDebug(t *testing.T) {
	root := CreateTree([]int{1, -1, 2, 2})
	res := findMode(root)
	fmt.Println(res)
}
