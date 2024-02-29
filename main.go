package main

import "fmt"

func main() {
	doAppend := func(s []int) {
		//fmt.Println(len(s), "   ", cap(s))
		s = append(s, 1)
		printLengthAndCapacity(s)
	}
	s := make([]int, 8, 8)
	fmt.Println(len(s[:4]), "   ", cap(s[:4]))

	doAppend(s[:4])
	printLengthAndCapacity(s)
	doAppend(s)
	printLengthAndCapacity(s)
	m := make(map[string]string)
	m["1"] = "1"
}

func printLengthAndCapacity(s []int) {
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d \n", len(s), cap(s))
}
