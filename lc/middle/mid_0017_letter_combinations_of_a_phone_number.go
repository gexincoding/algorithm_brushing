package middle

var phone = [][]rune{
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
	{'j', 'k', 'l'},
	{'m', 'n', 'o'},
	{'p', 'q', 'r', 's'},
	{'t', 'u', 'v'},
	{'w', 'x', 'y', 'z'},
}

var ans []string

func letterCombinations(digits string) []string {
	ans = make([]string, 0)
	dfs([]rune(digits), 0, make([]rune, 0))
	return ans
}

func dfs(digits []rune, index int, path []rune) {
	if index == len(digits) {
		return
	}
	dfs(digits, index, path)
}
