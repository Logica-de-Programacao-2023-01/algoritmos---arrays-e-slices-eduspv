package main

import "fmt"

func main() {
	array := [3][2]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Informe o valor do elemento [%d][%d]: ", i, j)
			fmt.Scan(&array[i][j])
		}
	}
	fmt.Println("Matriz resultante:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(array[i][j])
		}
	}
}
