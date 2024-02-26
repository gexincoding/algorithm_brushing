package _008_string_to_integer

import "math"

func myAtoi(s string) int {

	charInt := []rune(s)
	return int(getInt(charInt))
}

func getInt(charInt []rune) int32 {
	index := 0
	// 上面进行一些过滤条件
	neg := false
	if charInt[0] == '-' {
		neg = true
		index = 1
	}
	var res int32 = 0
	threshold1 := math.MinInt32 / 10
	threshold2 := math.MinInt32 % 10
	for ; index < len(charInt); index++ {
		// 弄成负数搞
		cur := '0' - charInt[index]
		// 如果会超范围，那么就溢出
		if res < int32(threshold1) || (res == int32(threshold1) && cur < int32(threshold2)) {
			if neg {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
		res = res*10 + cur
	}
	if !neg && res == math.MinInt32 {
		return math.MaxInt32
	}
	if neg {
		return res
	} else {
		return -res
	}
}
