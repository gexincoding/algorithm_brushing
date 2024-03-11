package _547_number_of_provinces

type UnionFind struct {
	father []int
	size   []int
	help   []int
	sets   int
}

func CreateUnionFind(n int) *UnionFind {
	unionFind := UnionFind{
		father: make([]int, n),
		size:   make([]int, n),
		help:   make([]int, n),
		sets:   n,
	}
	for i := 0; i < n; i++ {
		unionFind.father[i] = i
		unionFind.size[i] = i
	}
	return &unionFind

}

func (unionFind *UnionFind) Union(i, j int) {
	fatherI := unionFind.findFather(i)
	fatherJ := unionFind.findFather(j)
	if fatherJ == fatherI {
		return
	}
	if unionFind.size[fatherI] > unionFind.size[fatherJ] {
		unionFind.father[j] = i
		unionFind.size[i] = fatherI + fatherJ
	} else {
		unionFind.father[i] = j
		unionFind.size[j] = fatherI + fatherJ
	}
	unionFind.sets--
}
func (unionFind *UnionFind) Sets() int {
	return unionFind.sets
}

func (unionFind *UnionFind) findFather(i int) int {
	hi := 0
	for i != unionFind.father[i] {
		unionFind.help[hi] = i
		hi++
		i = unionFind.father[i]
	}
	for hi--; hi >= 0; hi-- {
		unionFind.father[unionFind.help[hi]] = i
	}
	return i
}
