package main

import (
	"fmt"
)

func main() {

	var hora, min, sec int
	var ampm string

	// Capturo o input da hora ja separando cada um em sua variavel
	fmt.Scanf("%d:%d:%d%s", &hora, &min, &sec, &ampm)

	// Operacao logica para identificar se eu acrescento horas ou nao
	if ampm == "AM" && hora == 12 {
		hora = 0
	} else if ampm == "PM" && hora != 12 {
		hora += 12
	}
	fmt.Printf("%02d:%02d:%02d", hora, min, sec)
}
