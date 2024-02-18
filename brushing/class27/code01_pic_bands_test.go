package class27

import (
	"fmt"
	"math"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := [][]int{
		{1, 2, 3},
		{4, 5},
	}
	changeSlice(arr)
	fmt.Println(arr)
}

func changeSlice(arr [][]int) {
	arr[0][1] = math.MaxInt32
}

func TestClean(t *testing.T) {
	programs := [][]int{
		{1, 2, 10},
		{2, 1, 30},
		{2, 1, 40},
		{1, 3, 20},
		{4, 5, 10},
	}
	size := clean(programs)
	fmt.Printf("the size is %d\n", size)

	for i := 0; i < size; i++ {
		fmt.Println(programs[i])
	}

}
