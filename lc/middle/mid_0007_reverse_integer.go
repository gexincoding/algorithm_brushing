package middle

import "math"

func reverse(x int) int {
	isNeg := x < 0
	// x 变成负数处理
	if !isNeg {
		x = -x
	}
	res := 0
	for x != 0 {
		if math.MinInt32/10 > res || (math.MinInt32/10 == res && math.MinInt32%10 > x%10) {
			return 0
		}
		res = res*10 + x%10
		x = x / 10
	}
	if !isNeg {
		return -res
	}
	return res
}
