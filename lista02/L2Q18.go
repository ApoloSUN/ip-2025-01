package main

import (
	"fmt"
	"strings"
)

func main() {
	var categoria, dia string 
	var precoN float64
	
	fmt.Println("Digite o preço normal, dia do aluguel e a categoria do DVD: ")
	fmt.Scanln(&precoN, &dia, &categoria)

	if ((dia == "2") || (dia == "3") || (dia == "5")) && (strings.ToUpper (categoria) == "COMUM") {
		precoF := precoN - (precoN * 0.4)
		fmt.Println("O preço final da locação é", precoF)
	} else if ((dia == "2") || (dia == "3") || (dia == "5")) && (strings.ToUpper (categoria) == "LANCAMENTO") {
		precoF := precoN - (precoN * 0.4) + (precoN * 0.15)
		fmt.Println("O preço final da locação é", precoF)
	} else if ((dia == "4") || (dia == "6") || (dia == "7") || (dia == "1")) && (strings.ToUpper (categoria) == "COMUM") {
		precoF := precoN
		fmt.Println("O preço final da locação é", precoF)
	} else if ((dia == "4") || (dia == "6") || (dia == "7") || (dia == "1")) && (strings.ToUpper (categoria) == "LANCAMENTO") {
		precoF := precoN + (precoN * 0.15)
		fmt.Println("O preço final da locação é", precoF)
	}
}