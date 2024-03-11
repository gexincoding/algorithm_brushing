package _547_number_of_provinces

import (
	"fmt"
	"testing"
)

func TestUnionFind(t *testing.T) {
	unionFind := CreateUnionFind(10)
	fmt.Println(unionFind.Sets())
	unionFind.Union(2, 3)
	fmt.Println(unionFind.Sets())
	unionFind.Union(2, 4)
	fmt.Println(unionFind.Sets())
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", unionFind.findFather(i))
	}
}
