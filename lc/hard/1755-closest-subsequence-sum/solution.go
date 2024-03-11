package _755_closest_subsequence_sum

import (
	"math"
	"sort"
)

func minAbsDifference(nums []int, goal int) int {
	res := math.MaxInt
	arrL := make([]int, 0)
	arrR := make([]int, 0)
	m := (len(nums) - 1) / 2
	f(nums, 0, 0, m, &arrL)
	f(nums, m+1, 0, len(nums)-1, &arrR)
	sort.Ints(arrL)
	sort.Ints(arrR)
	for i := 0; i < len(arrL); i++ {
		target := goal - arrL[i]
		for j := len(arrR) - 1; j >= 1; j-- {
			p1 := abs(arrR[j-1] - target)
			p2 := abs(arrR[j] - target)
			if p1 <= p2 {
				res = min(res, p1)
			} else {
				break
			}
		}
	}
	return res
}

func f(nums []int, index, sum, endIndex int, arr *[]int) {
	if index == endIndex+1 {
		*arr = append(*arr, sum)
		return
	}
	f(nums, index+1, sum, endIndex, arr)
	f(nums, index+1, sum+nums[index], endIndex, arr)
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
