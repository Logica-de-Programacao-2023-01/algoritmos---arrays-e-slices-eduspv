package main

import "fmt"

func main() {
	var soma float32
	numeros := [4]float32{1.1, 2.2, 3.3, 4.4}
	for _, x := range numeros {
		soma = soma + x
	}
	fmt.Println(soma)

}
