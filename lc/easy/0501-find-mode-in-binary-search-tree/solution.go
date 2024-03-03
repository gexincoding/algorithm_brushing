package _501_find_mode_in_binary_search_tree

import "math"

var (
	pre             *TreeNode
	res             []int
	historyMaxCount int
	curCount        int
)

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	pre = nil
	res = make([]int, 0)
	historyMaxCount = math.MinInt64
	dfs(root)
	return res
}

func dfs(cur *TreeNode) {
	if cur == nil {
		return
	}
	dfs(cur.Left)
	if pre == nil {
		curCount = 1
	} else if pre.Val == cur.Val {
		curCount++
	} else {
		curCount = 1
	}
	if curCount == historyMaxCount {
		res = append(res, cur.Val)
	}
	if curCount > historyMaxCount {
		res = make([]int, 0)
		historyMaxCount = curCount
		res = append(res, cur.Val)
	}

	pre = cur
	dfs(cur.Right)

}
