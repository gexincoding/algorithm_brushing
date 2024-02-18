package easy

func sortedSquares(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	size := len(nums)
	ans := make([]int, size)
	for i, j := 0, len(nums)-1; i <= j; {
		if nums[i]*nums[i] >= nums[j]*nums[j] {
			size--
			ans[size] = nums[i] * nums[i]
			i++
		} else {
			size--
			ans[size] = nums[j] * nums[j]
			j--
		}
	}
	return ans
}
