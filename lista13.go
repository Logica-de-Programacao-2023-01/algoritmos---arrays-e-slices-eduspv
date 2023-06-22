package main

import "fmt"

// Crie um Array de inteiros com 7 elementos.
// Solicite ao usuário que informe um número que será adicionado ao primeiro e ao último elemento do Array.
// Imprima o Array resultante.
func main() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}
	var n1, n2 int
	fmt.Print("informe um numero para ser adicionado no inicio da array:")
	fmt.Scan(&n1)
	fmt.Print("informe o segundo numero para ser adicionado na array")
	fmt.Scan(&n2)
	array[0] = n1
	array[6] = n2
	fmt.Print(array)
}
