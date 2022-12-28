package main

import (
	"fmt"
	"strconv"
)

func main() {
	var r uint32 = 1
	var s string = fmt.Sprintf("%032b", ^r)

	output, err := strconv.ParseUint(s, 2, 32)
	fmt.Println("Binario:", s, "\nConvertido:", output, err)
}
