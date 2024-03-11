package _746_min_cost_climbing_stairs

var cache map[int]int

func minCostClimbingStairs(cost []int) int {
	/*	cache = make(map[int]int)
		return getMin(f(0, cost), f(1, cost))*/

	dp := make([]int, len(cost)+2)
	dp[len(cost)], dp[len(cost)+1] = 0, 0
	for i := len(dp) - 3; i >= 0; i-- {
		dp[i] = getMin(dp[i+1]+cost[i], dp[i+2]+cost[i])
	}
	return getMin(dp[0], dp[1])

}

// 目前在index位置，想跳到顶楼 最小花费是多少？
// 已经花费了done
func f(index int, cost []int) int {
	ansCache := get(index)
	if ansCache != -1 {
		return ansCache
	}
	if index >= len(cost) {
		return 0
	}
	// 1. 跳一步
	p1 := f(index+1, cost) + cost[index]
	// 2. 跳两部
	p2 := f(index+2, cost) + cost[index]
	ans := getMin(p1, p2)
	set(index, ans)
	return ans
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func get(k int) int {
	if val, ok := cache[k]; ok {
		return val
	}
	return -1
}

func set(k, v int) {
	cache[k] = v
}
