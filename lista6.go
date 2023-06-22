package main

import (
	"fmt"
)

//Crie um Array de inteiros com 10 elementos. Em seguida, solicite ao usuário que informe um valor e verifique
//se esse valor está presente no Array. Imprima o resultado da busca.

func main() {
	var numero_digitado int
	numeros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultado := false
	fmt.Print("que numero voce acha que esta dentro dos conjuntos?")
	fmt.Scan(&numero_digitado)

	for _, x := range numeros {
		if numero_digitado == x {
			resultado = true
			break
		}
	}
	if resultado == true {
		fmt.Print("o seu numero esta na lista")
	} else {
		fmt.Print("o seu numero nao esta na lista.")
	}

}
