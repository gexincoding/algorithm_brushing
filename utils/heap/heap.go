package heap

type IntHeap []int

func (intHeap *IntHeap) Len() int {
	return len(([]int)(*intHeap))
}

func (intHeap *IntHeap) Less(i, j int) bool {
	return false
}
