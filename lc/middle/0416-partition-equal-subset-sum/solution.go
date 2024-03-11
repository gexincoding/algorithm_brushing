package _416_partition_equal_subset_sum

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum&1 == 1 {
		return false
	}
	tarSum := sum / 2
	dp := make([][]bool, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, 100*len(nums))
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = true
	}
	for i := 1; i < len(dp[0]); i++ {
		dp[len(dp)-1][i] = false
	}

	for index := len(dp) - 2; index >= 0; index-- {
		for tar := 1; tar < len(dp[0]); tar++ {
			p1 := false
			if tar-nums[index] >= 0 {
				p1 = dp[index+1][tar-nums[index]]
			}
			p2 := dp[index+1][tar]
			dp[index][tar] = p1 || p2
		}
	}
	return dp[0][tarSum]
}

// 当前来到index位置，num[0,index-1]范围已经做好选择
// 我需要求 index往后 能不能选出数的和是tar
func dfs(index int, nums []int, tar int) bool {
	if tar == 0 {
		return true
	}
	if index == len(nums) {
		return false
	}

	p1 := dfs(index+1, nums, tar-nums[index])
	p2 := dfs(index+1, nums, tar)
	return p1 || p2
}
