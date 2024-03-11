package _912_sort_an_array

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	array := sortArray([]int{1, 3, 2, 1, 24, 1, 24})
	for _, item := range array {
		fmt.Println(item)
	}
}
