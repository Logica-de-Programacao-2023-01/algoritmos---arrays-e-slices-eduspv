package main

import "fmt"

//Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que informe um número inteiro.
//Adicione esse número ao Slice apenas se ele não estiver presente. Imprima o Slice resultante.

func main() {
	add := 0
	numeros := []int{1, 2, 3, 4, 5}
	resultado := false

	fmt.Print("digite um numero para ser adicionado a slice:")
	fmt.Scan(&add)
	for _, x := range numeros {
		if add == x {
			resultado = true
			break
		}
	}
	if resultado == true {
		fmt.Print("o numero digitado ja existe na slice.")

	} else {
		numeros = append(numeros, add)
		fmt.Print(numeros)

	}

}
