package _062_unique_paths

func uniquePaths(m int, n int) int {
	/*return f(0, 0, m-1, n-1)*/
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[i][j] = 1
				continue
			}
			ans := 0
			if i+1 <= m-1 {
				ans += dp[i+1][j]
			}
			if j+1 <= n-1 {
				ans += dp[i][j+1]
			}
			dp[i][j] = ans
		}
	}
	return dp[0][0]
}

func f(i, j, tarX, tarY int) int {
	if i > tarX || j > tarY {
		return -1
	}
	if i == tarX && j == tarY {
		return 1
	}
	ans := 0
	p1 := f(i+1, j, tarX, tarY)
	p2 := f(i, j+1, tarX, tarY)
	if p1 != -1 {
		ans += p1
	}
	if p2 != -1 {
		ans += p2
	}
	return ans
}
