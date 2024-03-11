package _063_unique_paths_ii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 1 && len(obstacleGrid[0]) == 1 {
		if obstacleGrid[0][0] == 0 {
			return 1
		}
		return 0

	}
	/*	return f(0, 0, obstacleGrid)*/
	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	for i := len(obstacleGrid) - 1; i >= 0; i-- {
		for j := len(obstacleGrid[0]) - 1; j >= 0; j-- {
			if i == len(obstacleGrid)-1 && j == len(obstacleGrid[0])-1 {
				dp[i][j] = 1
				continue
			}
			ans := 0
			if i+1 <= len(obstacleGrid)-1 && obstacleGrid[i+1][j] != 1 {
				ans += dp[i+1][j]
			}
			if j+1 <= len(obstacleGrid[0])-1 && obstacleGrid[i][j+1] != 1 {
				ans += dp[i][j+1]
			}
			dp[i][j] = ans
		}
	}
	return dp[0][0]
}

// 当前的位置是 (x,y) 走到 (len(obstacleGrid)-1, len(obstacleGrid[0])-1) 有多少种走法
func f(x, y int, obstacleGrid [][]int) int {
	if x == len(obstacleGrid)-1 && y == len(obstacleGrid[0])-1 {
		return 1
	}
	if x > len(obstacleGrid)-1 || y > len(obstacleGrid[0])-1 {
		return 0
	}
	if obstacleGrid[x][y] == 1 {
		return 0
	}
	p1 := f(x+1, y, obstacleGrid)
	p2 := f(x, y+1, obstacleGrid)
	return p1 + p2
}
