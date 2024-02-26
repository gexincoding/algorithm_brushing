package class01

import (
	"math"
	"testing"
)

func TestCoverMaxPoint(t *testing.T) {
	/*
	   len := 100
	   max := 1000
	   testTime := 100000
	   fmt.Println("测试开始")

	   	for i := 0; i < testTime; i++ {
	   		L := rand.Intn(max)
	   		arr := generateArray(len, max)
	   		ans1 := maxPoint1(arr, L)
	   		ans2 := maxPoint2(arr, L)
	   		ans3 := utils(arr, L)
	   		if ans1 != ans2 || ans2 != ans3 {
	   			fmt.Println("oops!")
	   			break
	   		}
	   	}
	*/
}

func test(arr []int, L int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		pre := i - 1
		for pre >= 0 && arr[i]-arr[pre] <= L {
			pre--
		}
		max = int(math.Max(float64(max), float64(i-pre)))
	}
	return max
}
