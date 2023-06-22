package main

//Crie um programa que leia um array de inteiros e verifique se ele está ordenado em ordem crescente.

import (
	"fmt"
)

func main() {
	array := [5]int{1, 3, 5, 7, 9}

	ordenado := true
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			ordenado = false
			break
		}
	}

	if ordenado {
		fmt.Println("O array está ordenado em ordem crescente.")
	} else {
		fmt.Println("O array não está ordenado em ordem crescente.")
	}
}
