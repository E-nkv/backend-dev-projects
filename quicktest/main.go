package main

import (
	"fmt"
	"slices"
)

func main() {
	asd := []int{1, 2, 3, 4, 5, 6}
	res := slices.DeleteFunc(asd, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(asd, len(asd), cap(asd))
	fmt.Println(res, len(res), cap(res))
	fmt.Println(&asd[1], &res[1])
}
