package main

import "fmt"

//Crie um Array de inteiros com 5 elementos.
//Em seguida, crie um novo Slice que contenha apenas os elementos do Array que são múltiplos de 3.
//Imprima o novo Slice.

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{}
	for _, x := range array {
		if x%2 > 0 {
			slice = append(slice, x)
		}

	}
	fmt.Print(slice)

}
