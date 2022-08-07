package main

import (
	"fmt"
)

func main() {

	// Crio as variaveis e capturo o tamanho do array
	var input, r int
	fmt.Scan(&input)

	// Crio o array com o tamanho desejado do input
	arr := make([]int, input)

	// Preencho o array
	for i := 0; i < input; i++ {
		fmt.Scan(&arr[i])
	}

	// Logica que compara se o numero se repete no array
	for i := 0; i < len(arr); i++ {
		r = 0
		for j := 0; j < len(arr); j++ {
			if arr[j] == arr[i] && i != j {
				r++
			}
		}
		if r == 0 {
			fmt.Println(arr[i])
		}
	}
}
