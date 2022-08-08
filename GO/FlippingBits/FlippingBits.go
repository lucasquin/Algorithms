package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Capturo o tamanho do array
	var arrSize int
	fmt.Scan(&arrSize)

	// Crio o array com o tamanho desejado do input
	arr := make([]uint32, arrSize)

	// Preencho o array
	for i := 0; i < arrSize; i++ {
		fmt.Scan(&arr[i])
	}

	// Logica para flipar os bits
	for k := 0; k < arrSize; k++ {
		s := fmt.Sprintf("%032b", ^arr[k])
		output, err := strconv.ParseUint(s, 2, 32)
		fmt.Println(output)
		if err != nil {
			panic(err)
		}
	}
}
