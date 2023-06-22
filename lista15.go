package main

import "fmt"

func main() {
	array := [10]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := []float32{}
	maior := false
	for _, x := range array {
		if x > 5 {
			maior = true
		}
		if maior == true {
			slice = append(slice, x)
		}
	}
	fmt.Print(slice)

}
