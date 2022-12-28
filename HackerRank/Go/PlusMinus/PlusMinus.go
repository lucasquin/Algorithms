package main

import "fmt"

func main() {

	// Declaro as variaveis que utilizarei
	var arr, input, pos, neg, z int

	// Obtenho o tamanho do array
	fmt.Scan(&arr)

	// Com o tamanho do array em maos, eu vou preenchendo ele...
	for i := 0; i < arr; i++ {

		// Capturo o valor de entrada de cada membro que vai compor o array
		fmt.Scan(&input)

		// Logica para diferenciar os numeros positivos, negativos e zeros do array
		if input > 0 {
			pos++
		} else if input < 0 {
			neg++
		} else {
			z++
		}
	}

	// Divisao feita para obter a proporcao entre os numeros positivos, negativos e zeros dentro do array
	positivo := float64(pos) / float64(arr)
	negativo := float64(neg) / float64(arr)
	zero := float64(z) / float64(arr)

	// Imprimo o resultado com 6 casas decimais de precicao
	fmt.Printf("%.6f\n", positivo)
	fmt.Printf("%.6f\n", negativo)
	fmt.Printf("%.6f\n", zero)
}
