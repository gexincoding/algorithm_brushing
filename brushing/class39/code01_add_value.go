package class39

// 来自腾讯
// 给定一个只由0和1组成的字符串S，假设下标从1开始，规定i位置的字符价值V[i]计算方式如下:
// 1) i == 1时，V[i] = 1
// 2) i > 1时，如果S[i] != S[i-1]，V[i] = 1
// 3) i > 1时，如果S[i] == S[i-1]，V[i] = V[i-1] + 1
// 你可以随意删除S中的字符，返回整个S的最大价值
// 字符串长度<=5000

func getMaxValue(s string) int {
	return f([]rune(s), 0, -1, 0)
}

// s 为字符穿
// index 为当前需要做决定的位置
// last 为左边离我最近的没删掉的位置的值是0还是1
// lastAns 为左边离我最近的没删掉的位置的价值是多少

func f(s []rune, index int, last rune, lastAns int) int {
	if index == len(s) {
		return lastAns
	}

	var chosenValue = lastAns
	if last == s[index] {
		chosenValue += lastAns + 1

	} else {
		chosenValue += 1
	}
	p1 := f(s, index+1, s[index], chosenValue)
	p2 := f(s, index+1, last, lastAns)
	if p1 > p2 {
		return p1
	}
	return p2

}
