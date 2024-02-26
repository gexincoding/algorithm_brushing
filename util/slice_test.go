package util

import (
	"fmt"
	"math"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := make([]int, 2, 1000)
	arr[0] = 1
	arr[1] = 10

	fmt.Printf("%p\n", arr)
	arr = append(arr, 9)
	fmt.Printf("%p\n", arr)
	fmt.Println(arr)
	//f(&arr)

}

func f(arr *[]int) {
	(*arr)[0] = math.MaxInt
	fmt.Printf("%p\n", arr)

	*arr = append(*arr, 100)
	fmt.Printf("%p\n", arr)
}
