package main

import "fmt"

//Crie um Slice de strings com tamanho 8 e solicite ao usuário que informe um valor.
//Remova todas as ocorrências desse valor no Slice e imprima o resultado.

func main() {
	original := []string{"a", "b", "a", "c", "a", "x", "i", "s"}
	var retirar string
	encontrou := false
	novaslice := []string{}
	fmt.Print("digite uma letra para ser retirada da palavra (abacaxi):")
	fmt.Scan(&retirar)
	for _, x := range original {
		if retirar != x {
			novaslice = append(novaslice, x)
		} else {
			encontrou = true
		}
	}
	if encontrou {
		fmt.Println("Slice original:", original)
		fmt.Println("Slice resultante:", novaslice)
	} else {
		fmt.Println("A letra digitada não está na slice")
	}
}
