package easy

func removeElement(nums []int, val int) int {
	size := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[size] = nums[i]
			size++
		}
	}
	return size
}
