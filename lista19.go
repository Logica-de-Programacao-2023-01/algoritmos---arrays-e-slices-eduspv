package main

import "fmt"

//Faça um programa que leia dois arrays de inteiros de tamanho n e gere um terceiro array que seja a soma dos dois primeiros.

func main() {
	array1 := [5]int{1, 2, 3, 4, 5}
	array2 := [3]int{1, 2, 3}
	array3 := make([]int, len(array1))
	for i := 0; i < len(array2); i++ {
		for t := 0; t < len(array1); t++ {

			array3[t] = array1[t] + array2[i]
		}
	}
	fmt.Println(array3)
}
