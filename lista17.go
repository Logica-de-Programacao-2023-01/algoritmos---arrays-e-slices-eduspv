package main

import "fmt"

func main() {

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum := 0

	for i := 0; i < len(arr); i += 2 {
		sum += arr[i]
	}

	fmt.Println("A soma dos elementos nas posições pares do array é:", sum)
}
