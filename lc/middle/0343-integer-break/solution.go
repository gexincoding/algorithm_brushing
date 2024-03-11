package _343_integer_break

func integerBreak(n int) int {
	return f(n)
}

// left 数字，拆分之后乘积最大值返回
func f(left int) int {
	if left == 0 {
		return 1
	}
	ans := 1
	for i := 1; i < left; i++ {
		p1 := f(left-i) * i
		p2 := (left - i) * i
		ans = getMax(getMax(p1, p2), ans)
	}
	return ans
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
