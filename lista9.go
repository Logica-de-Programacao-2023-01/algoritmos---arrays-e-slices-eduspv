package main

import "fmt"

//Crie um Array de floats com 6 elementos.
//Solicite ao usuário que informe um número que será adicionado em todas as posições do Array.
//Imprima o Array resultante.

func main() {
	elementos := make([]float32, 6)
	var add float32
	fmt.Print("informe um numero para ser adicionado em todas as posições da array:")
	fmt.Scan(&add)
	for i := 0; i < len(elementos); i++ {
		elementos[i] = add
	}
	fmt.Println("array resultante é", elementos)

}
