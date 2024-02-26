package _001_two_sum

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, item := range nums {
		index, ok := numMap[target-item]
		if ok {
			return []int{index, i}
		} else {
			numMap[item] = i
		}
	}
	return []int{-1, -1}
}
