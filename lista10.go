package main

import "fmt"

//Crie um Slice de inteiros com tamanho 10 e imprima o valor mínimo e o valor máximo armazenados no Slice.

func main() {
	slice := []int32{10, 2, 3, 14, 7, 6, 7, 15, 9, 11}
	var valor int32
	valor2 := 11

	for _, x := range slice {
		if x > valor {
			valor = x
		}
	}
	fmt.Println("o maior numero é:", valor)
	for _, y := range slice {
		if int(y) < valor2 {
			valor2 = int(y)
		}

	}
	fmt.Print("o menor numero é:", valor2)
}
