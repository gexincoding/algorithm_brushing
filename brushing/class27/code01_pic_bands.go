package class27

import (
	"math"
	"sort"
)

//每一个项目都有三个数，[a,b,c]表示这个项目a和b乐队参演，花费为c
//每一个乐队可能在多个项目里都出现了，但是只能被挑一次
//nums是可以挑选的项目数量，所以一定会有nums*2只乐队被挑选出来
//返回一共挑nums轮(也就意味着一定请到所有的乐队)，最少花费是多少？
//如果怎么都无法在nums轮请到nums*2只乐队且每只乐队只能被挑一次，返回-1
//nums<9，programs长度小于500，每组测试乐队的全部数量一定是nums*2，且标号一定是0 ~ nums*2-1

func getMinCost(programs [][]int, nums int) int {
	// 创建一个map
	// 选中状态与最小代价之间的映射
	// 下标代表一种选中的状态
	// 由于一共有nums * 2只乐队，因此需要nums * 2 位来放置状态
	//minCostMap = make([]int, (1<<(nums<<1))-1)
	size := clean(programs)
	var minCostMap1 = make([]int, 1<<(nums<<1))
	initArr(minCostMap1)
	var minCostMap2 []int = nil
	if nums&1 == 0 { // 如果nums是偶数，可以平分nums进行分治
		dfs(programs, size, 0, nums>>1, 0, 0, minCostMap1)
		minCostMap2 = minCostMap1
	} else {
		dfs(programs, size, 0, nums>>1, 0, 0, minCostMap1)
		minCostMap2 = make([]int, 1<<(nums<<1))
		initArr(minCostMap2)
		dfs(programs, size, 0, nums-(nums>>1), 0, 0, minCostMap2)
	}
	minCost := math.MaxInt32
	mask := (1 << (nums << 1)) - 1 // 为了砍掉 取反后一大堆高位的1 （通过和mask进行按位与运算）
	for i := 0; i < len(minCostMap1); i++ {
		if minCostMap1[i] != math.MaxInt32 && minCostMap2[mask&(^i)] != math.MaxInt32 {
			minCost = getMin(minCost, minCostMap1[i]+minCostMap2[mask&(^i)])
		}
	}
	if minCost == math.MinInt32 {
		minCost = -1
	}
	return minCost
}
func initArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = math.MaxInt32
	}
}

// index 当前决策第index号的项目
// rest 剩余还能挑选多少项目 （每挑选一个项目 相当于是 挑选了两个乐队）
// pick 选了哪些乐队（用一个整数的二进制信息）
// done  所有乐队都挑到的状态（乐队个数为2*nums）：32位的整数，低2*nums位全为1 ，标示为：(1 << 2 * nums) - 1 下面这种写法是直接把结果算出写到参数里，原因是和nums解耦
// 2 * nums 也能写成 nums << 1 因此能够写成：(1 << (nums << 1)) - 1
// cost 0～index的总花费
func dfs(programs [][]int, size int, index int, rest int, pick int, cost int, minCostMap []int) {
	if rest == 0 { // 没有项目可挑
		//minCost = getMin(minCost, cost)
		minCostMap[pick] = getMin(minCostMap[pick], cost)
	} else { // 还有项目可挑
		if index != size {
			// 不考虑当前的项目
			dfs(programs, size, index+1, rest, pick, cost, minCostMap)
			// 考虑当前的项目
			pickA := 1 << programs[index][0]
			pickB := 1 << programs[index][1]
			cur := pickA | pickB
			// 这两只乐队 之前没被挑过 才能被挑
			if cur&pick == 0 {
				dfs(programs, size, index+1, rest-1, pick|cur, cost+programs[index][2], minCostMap)
			}
		}
	}
}

func clean(programs [][]int) int {

	// 小数组内部排序
	for _, subArr := range programs {
		if subArr[0] > subArr[1] {
			temp := subArr[0]
			subArr[0] = subArr[1]
			subArr[1] = temp
		}
	}

	sort.Slice(programs, func(i, j int) bool {
		if programs[i][0] == programs[j][0] {
			if programs[i][1] == programs[j][1] {
				return programs[i][2] < programs[j][2]
			} else {
				return programs[i][1] < programs[j][1]

			}
		} else {
			return programs[i][0] < programs[j][0]
		}
	})

	size := 1
	// 根据x和y唯一确定一组
	// 起始组
	x := programs[0][0]
	y := programs[0][1]

	// 从第二个字数组开始
	for i := 1; i < len(programs); i++ {
		// 如果还是这一组，就不管，跳过，如果不是这一组，继续
		if x != programs[i][0] && y != programs[i][1] { // 说明遇到了下一组
			// 记录下一组
			programs[size] = programs[i]
			x = programs[i][0]
			y = programs[i][1]
			size++
		}
	}
	return size
}

func getMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))

}

func getMax(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}
