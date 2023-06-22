package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&n)

	primes := make([]int, n)
	count := 0
	num := 2

	for count < n {
		prime := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				prime = false
				break
			}
		}
		if prime {
			primes[count] = num
			count++
		}
		num++
	}

	fmt.Println(primes)
}
