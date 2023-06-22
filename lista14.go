package main

import "fmt"

//Crie um Slice de inteiros com tamanho 8 e solicite ao usuário que informe dois índices
//de elementos que deverão ser trocados de posição.
//Imprima o Slice resultante.

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var i1, i2 int
	fmt.Print("informe o primeiro indice(entre 0 e 7):")
	fmt.Scan(&i1)
	fmt.Print("informe o segundo indice(entre 0 e 7):")
	fmt.Scan(&i2)

	i1 = i1 % len(slice)
	i2 = i2 % len(slice)
	slice[i1], slice[i2] = slice[i2], slice[i1]
	fmt.Print(slice)

}
