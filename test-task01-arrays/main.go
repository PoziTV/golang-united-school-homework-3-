package main

import (
	"fmt"
	"homework"
	// "../homework"
)

func main() {
	slice1 := make([]int64, 0)
	slice1 = append(slice1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)

	out := homework.Reverse(slice1)
	fmt.Println(out)

	slice2 := make([]int64, 0)
	slice2 = append(slice2, 3, 5, 2, 8)
	out2 := homework.Reverse(slice2)
	fmt.Println(out2)

}
