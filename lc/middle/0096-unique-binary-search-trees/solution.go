package _096_unique_binary_search_trees

func numTrees(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	/*return dfs(n)*/
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2
	for x := 3; x <= n; x++ {
		ans := 0
		for i := 1; i <= x; i++ {
			p1 := dp[i-1]
			p2 := dp[x-i] // 右树的组合数
			ans += p1 * p2
		}
		dp[x] = ans
	}
	return dp[n]
}

// 节点的取值为1～x，返回搜索二叉树的数量
func dfs(x int) int {
	if x == 0 || x == 1 {
		return 1
	}
	if x == 2 {
		return 2
	}
	// 枚举头节点是谁
	// i为选中的头节点
	ans := 0

	for i := 1; i <= x; i++ {
		p1 := dfs(i - 1) // 左树的组合数
		p2 := dfs(x - i) // 右树的组合数
		ans += p1 * p2
	}
	return ans
}
