package util

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"unsafe"
)

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	s = append(s[:1], s[2:]...)
	//f(&arr)

}

func f(arr *[]int) {
	(*arr)[0] = math.MaxInt
	fmt.Printf("%p\n", arr)

	*arr = append(*arr, 100)
	fmt.Printf("%p\n", arr)
}

func TestGolang(t *testing.T) {
	s := make([]int, 5, 10)
	PrintSliceStruct(&s)
	test(s)

}

func test(s []int) {
	PrintSliceStruct(&s)

}
func PrintSliceStruct(s *[]int) {
	// 代码 将slice转换成 reflect.SliceHeader
	ss := (*reflect.SliceHeader)(unsafe.Pointer(s))
	//查看slice的结构
	fmt.Printf("slice struct: %+v, slice is %v\n", ss, s)
}
