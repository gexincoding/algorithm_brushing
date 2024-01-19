package middle

import (
	"fmt"
	"testing"
)

func TestRabbit(t *testing.T) {
	answers := []int{1, 1, 2}
	fmt.Println(numRabbits(answers))

	answers2 := []int{10, 10, 10}
	fmt.Println(numRabbits(answers2))

	ansWrong := []int{1, 0, 1, 0, 0}
	fmt.Println(numRabbits(ansWrong))

}
