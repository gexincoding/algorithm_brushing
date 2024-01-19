package middle

func numRabbits(answers []int) int {
	ans := 0
	rabbitSet := make(map[int]interface{})
	for _, item := range answers {
		_, ok := rabbitSet[item+1]
		if !ok { // 存在这个记录
			rabbitSet[item+1] = nil
		}
	}
	if len(rabbitSet) == 0 {
		return -1
	} else {
		for k, _ := range rabbitSet {
			ans += k
		}
	}
	return ans
}
