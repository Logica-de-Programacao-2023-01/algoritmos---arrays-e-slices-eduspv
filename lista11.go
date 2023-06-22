package main

import "fmt"

//Crie um Array bidimensional de inteiros com 2 linhas e 3 colunas.
//Em seguida, solicite ao usuário que informe um índice de linha e
//outro de coluna e imprima o valor armazenado nessa posição
//da matriz.

func main() {
	matriz := [2][3]int{{1, 2, 5}, {3, 4, 6}}
	var linha, coluna int
	fmt.Println(matriz)
	fmt.Print("Digite um índice de linha (0 ou 1): ")
	fmt.Scanln(&linha)
	fmt.Print("Digite um índice de coluna (0, 1 ou 2): ")
	fmt.Scanln(&coluna)

	// Imprimindo o valor armazenado na posição informada pelo usuário
	fmt.Printf("O valor armazenado na posição (%d, %d) da matriz é o: %d\n", linha, coluna, matriz[linha][coluna])

}
