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
		index := find(arrR, 0, len(arrR)-1, target)
		if index == -1 {
			res = min(res, abs(arrR[0]-target))
		} else if index == len(arrR)-1 {
			res = min(res, abs(arrR[index]-target))
		} else {
			res = min(res, min(abs(arrR[index]-target), abs(arrR[index+1]-target)))
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

func find(arr []int, l, r int, tar int) int {
	index := -1
	for l <= r {
		m := (l + r) / 2
		if arr[m] <= tar {
			index = m
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return index
}
