// Link do desafio: https://www.hackerrank.com/challenges/sparse-arrays/problem

package main

import (
	"fmt"
)

func main() {

	var arrSize, r int

	// Capturo o tamanho do array
	fmt.Scan(&arrSize)

	// Definindo o tamanhdo do array de string
	arrString := make([]string, arrSize)

	// Inserindo valores nos indices do array de string
	for i := 0; i < arrSize; i++ {
		fmt.Scan(&arrString[i])
	}

	// Capturo o tamanho do array
	fmt.Scan(&arrSize)

	// Definindo o tamanhdo do array de query
	arrQuery := make([]string, arrSize)

	// Inserindo valores nos indices do array de query
	for i := 0; i < arrSize; i++ {
		fmt.Scan(&arrQuery[i])
	}

	// Logica para comparar quantos indices do array de query se repetem no array de string
	for i := 0; i < len(arrQuery); i++ {
		for j := 0; j < len(arrString); j++ {
			if arrString[j] == arrQuery[i] {
				r++
			}
		}
		// Imprimo o resultado da contagemde cada indice do array de query
		fmt.Println(r)
		// Zero o resultado para ir para o proximo
		r = 0
	}
}
