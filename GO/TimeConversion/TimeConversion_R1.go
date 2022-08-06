// Link do desafio: https://www.hackerrank.com/challenges/time-conversion/problem

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// Capturo o input da hora
	var horaRecebida string
	fmt.Scan(&horaRecebida)

	// Crio um array separando horas, minutos e segundos
	var arr = strings.Split(horaRecebida, ":")
	var hora, err = strconv.Atoi(arr[0])
	var minuto = arr[1]
	var segundo = strings.Split(arr[2], "")

	// Logica para verificar se e AM ou PM
	if strings.Contains(horaRecebida, "PM") {

		// Adiciono mais 12 horas
		if arr[0] != "12" {
			hora = hora + 12
		}
		// Crio o output para PM
		fmt.Printf("%v:%v:%v", hora, minuto, segundo[0]+segundo[1])
		fmt.Println()

	} else if strings.Contains(horaRecebida, "AM") {
		if arr[0] == "12" {
			hora = hora + 12
			if hora == 24 {
				hora = 0
			}
		}
		fmt.Printf("%02v:%v:%v", hora, minuto, segundo[0]+segundo[1])
		fmt.Println()

	} else {
		fmt.Println(err)
	}
}
