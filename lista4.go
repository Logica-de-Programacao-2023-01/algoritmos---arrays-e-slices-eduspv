package main

import "fmt"

func main() {
	var size int
	fmt.Print("Informe o tamanho do Slice: ")
	fmt.Scanln(&size)

	slice := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Printf("Informe o valor para o elemento %d :", i+1)
		fmt.Scanln(&slice[i])
	}

	fmt.Println("Slice: ", slice)

	soma := 0
	for _, value := range slice {
		soma += value
	}

	fmt.Println("Soma dos valores: ", soma)
}
