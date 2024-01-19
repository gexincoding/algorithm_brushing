package class01

import "math"

// 给定一个有序数组arr,代表坐落在X轴上的点
// 给定一个正数K,代表绳子的长度
// 返回绳子最多压中几个点?
// 即使绳子边缘处盖住点也算盖住
func maxPoint1(arr []int, L int) int {
	// 结尾处必定压住一个点，因此答案最小是1
	var res = 1

	for i := 0; i < len(arr); i++ {
		// 以i位置结尾
		nearest := nearestIndex(arr, i, arr[i]-L)
		res = int(math.Max(float64(res), float64(i-nearest+1)))
	}
	return res
}

func nearestIndex(arr []int, i int, x int) int {
	return 0
}
