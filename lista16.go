package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := []int{}

	for _, x := range array {
		if x%2 == 0 {

			slice = append(slice, x)
		}

	}
	fmt.Print(slice)
}
