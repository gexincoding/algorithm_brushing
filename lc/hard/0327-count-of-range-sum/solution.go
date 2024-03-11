package _327_count_of_range_sum

func countRangeSum(nums []int, lower int, upper int) int {
	return 0

}

var ans = 0

// 求 L～R 范围上 有组的每一个数，在左组有多少个数在该数所计算出的范围上！
// 满足条件的总数返回
// 如果L == R 只需要看 preSums[L] 是否在原始范围上，如果在，也达标！
func f(preSums []int, L, R int, lower, upper int) int {
	if L == R {
		if preSums[L] <= upper && preSums[L] >= lower {
			return 1
		}
		return 0
	}
	if L > R {
		return 0
	}
	M := (L + R) / 2
	p1 := f(preSums, L, M, lower, upper)
	p2 := f(preSums, M+1, R, lower, upper)
	p3 := merge(preSums, L, R, M, lower, upper)
	return p1 + p2 + p3
}

func merge(preSums []int, L, R, M int, lower, upper int) int {
	/*windowL, windowR := L, L
	for i := M + 1; i <= R; i++ {
		maxLen := 0
		// 当前寻找右组中的i位置 左组中有多少数达标！
		// 先计算出指标
		r := preSums[i] - lower
		l := preSums[i] - upper

		for windowL >= M+1 && windowL <= R && preSums[windowL] {
			windowL++
		}
		for windowR < r {

		}
	}*/
	return 0
}

func getMax(a, b int) int {

	if a > b {
		return a
	}
	return b
}
