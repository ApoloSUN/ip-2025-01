package main

import (
	"fmt"
	"strings"
)

func main() {
	var conta int64
	var tipo string 
	var consumo, preco float64
	
	fmt.Println("Digite a conta, o tipo de consumido (R/C/I) e o consumo de água: ")
	fmt.Scanln(&conta, &tipo, &consumo)

	if strings.ToUpper (tipo) == "R" {
		preco = 5 + (0.05 * consumo)
		fmt.Println(conta, "R$", preco)
	} else if strings.ToUpper (tipo) == "C" {
		preco = 500 + (0.25 * (consumo - 80))
		fmt.Println(conta, "R$", preco)
	} else {
		preco = 800 + (0.04 * (consumo - 100))
		fmt.Println(conta, "R$", preco)
	}
}