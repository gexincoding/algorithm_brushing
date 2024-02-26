package _476_closest_nodes_queries_in_a_binary_search_tree

import "math"

func closestNodes(root *TreeNode, queries []int) [][]int {
	var mini int
	var maxi int
	res := make([][]int, 0)
	for _, query := range queries {
		mini = math.MaxInt32
		maxi = math.MinInt32
		cur := root
		for cur != nil {
			if cur.Val == query {
				mini = cur.Val
				maxi = cur.Val
				break
			}
			if cur.Val > query {
				mini = getMin(mini, cur.Val)
				cur = cur.Left
			} else {
				cur = cur.Right
				maxi = getMax(cur.Val, maxi)

			}
		}

		if maxi == math.MinInt32 {
			maxi = -1
		}
		if mini == math.MaxInt32 {
			mini = -1
		}

		res = append(res, []int{maxi, mini})
	}
	return res
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func getMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}
