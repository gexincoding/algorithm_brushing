package _704_binary_search

// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，
// 写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := l + ((r - l) >> 1)
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}
