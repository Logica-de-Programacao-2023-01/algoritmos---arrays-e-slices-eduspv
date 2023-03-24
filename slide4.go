package main

import "fmt"

func main() {
	var somatorio float32
	numeros := [6]float32{2.2, 3.3, 4.4, 5.5, 6.6, 7.7}
	for _, x := range numeros {
		somatorio = somatorio + x

	}
	media := somatorio / 6
	fmt.Println(media)
}
