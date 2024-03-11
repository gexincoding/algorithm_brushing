package _547_number_of_provinces

func findCircleNum(isConnected [][]int) int {
	unionFind := CreateUnionFind(len(isConnected))
	for i := 0; i <= len(isConnected)-2; i++ {
		for j := i + 1; j <= len(isConnected[0])-1; j++ {
			if isConnected[i][j] == 1 {
				unionFind.Union(i, j)
			}
		}
	}
	return unionFind.Sets()
	return 0
}
