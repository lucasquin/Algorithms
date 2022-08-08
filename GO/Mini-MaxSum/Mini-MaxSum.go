package main

import (
	"fmt"
	"sort"
)

func main() {

	// Declaro o tamanho do array e preencho ele com meu input
	var arr [5]int
	var min int
	var max int

	fmt.Scan(&arr[0], &arr[1], &arr[2], &arr[3], &arr[4])

	// Ordeno o array em ordem crescente
	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i] < arr[j]
	})

	// Somo os 4 primeiros numeros menores
	for i := 0; i < 4; i++ {
		min = min + arr[i]
	}

	// Ordeno o array em ordem decrescente
	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i] > arr[j]
	})

	// Somo os 4 primeiros numeros maiores
	for i := 0; i < 4; i++ {
		max = max + arr[i]
	}

	// Imprimo
	fmt.Println(min, max)
}
