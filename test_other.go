package main

import (
	"fmt"
)

func main() {
	s1 := []int{0, 1, 2, 3}
	s2 := []int{4, 5, 6, 7}
	s1 = append(s1, s2...) // instead of FOR
	fmt.Println(s1)
}
