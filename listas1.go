package main

import "fmt"

func main() {
	var soma int
	numeros := [3]int{5, 2, 3}
	for _, x := range numeros {

		soma = soma + x

	}
	fmt.Print(soma)

}
