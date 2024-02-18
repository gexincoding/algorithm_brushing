package middle

import "sort"

//森林中有未知数量的兔子
//提问其中若干只兔子 "还有多少只兔子与你（指被提问的兔子）颜色相同?" ，
//将答案收集到一个整数数组 answers 中，其中 answers[i] 是第 i 只兔子的回答。
//给你数组 answers ，返回森林中兔子的最少数量。
// 1 <= answers.length <= 1000
// 0 <= answers[i] < 1000

func numRabbits(answers []int) int {

	sort.Slice(answers, func(i, j int) bool {
		return answers[i] < answers[j]
	})

	if len(answers) == 1 {
		return answers[0] + 1
	}

	curTermSum := 1
	ans := 0
	for index := 1; index < len(answers); index++ {
		if answers[index] != answers[index-1] {
			// collect answer
			num := answers[index-1]
			ans += ((curTermSum + num) / (num + 1)) * (num + 1)
			curTermSum = 1
			// a / b 向上取整 简便算法： (a + b - 1) / b
		} else {
			curTermSum++
		}
	}
	return ans + ((curTermSum+answers[len(answers)-1])/(answers[len(answers)-1]+1))*(answers[len(answers)-1]+1)
}
