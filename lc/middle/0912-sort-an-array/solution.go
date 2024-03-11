package _912_sort_an_array

func sortArray(nums []int) []int {
	mergeSort(&nums, 0, len(nums)-1)
	return nums
}

func mergeSort(arr *[]int, L, R int) {
	if L == R {
		return
	}
	M := (R + L) / 2
	mergeSort(arr, L, M)
	mergeSort(arr, M+1, R)
	merge(arr, L, R, M)
}

func merge(arr *[]int, L, R, M int) {
	l := L
	r := M + 1
	help := make([]int, R-L+1)
	i := 0
	for l <= M && r <= R {
		if (*arr)[l] <= (*arr)[r] {
			help[i] = (*arr)[l]
			l++
			i++
		} else {
			help[i] = (*arr)[r]
			r++
			i++
		}
	}
	for l <= M {
		help[i] = (*arr)[l]
		l++
		i++
	}
	for r <= R {
		help[i] = (*arr)[r]
		i++
		r++
	}
	for index, val := range help {
		(*arr)[L+index] = val
	}

}
